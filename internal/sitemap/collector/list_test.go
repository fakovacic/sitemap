package collector_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/fakovacic/sitemap/internal/sitemap/collector"
	"github.com/matryer/is"
)

func TestList(t *testing.T) {
	cases := []struct {
		it    string
		pages []sitemap.Page

		expectedResult []sitemap.Page
	}{
		{
			it: "it return all empty pages",
			pages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			expectedResult: []sitemap.Page{},
		},
		{
			it: "it return all ok pages",
			pages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
						Path:   "test",
					},
					Ok: true,
				},
			},
			expectedResult: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
						Path:   "test",
					},
					Ok: true,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			s := collector.New()
			ctx := context.TODO()

			for i := range tc.pages {
				_ = s.Add(ctx, tc.pages[i])
			}

			list := s.List(ctx)

			checkIs.Equal(list, tc.expectedResult)
		})
	}
}
