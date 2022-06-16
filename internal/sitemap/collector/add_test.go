package collector_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/fakovacic/sitemap/internal/sitemap/collector"
	"github.com/matryer/is"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		it    string
		pages []sitemap.Page

		expectedResult []bool
	}{
		{
			it: "it add page to collector",
			pages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			expectedResult: []bool{
				true,
			},
		},
		{
			it: "it add multiple pages to collector",
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
						Path:   "testing",
					},
				},
			},
			expectedResult: []bool{
				true,
				true,
			},
		},
		{
			it: "it add multiple same pages to collector",
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
					},
				},
			},
			expectedResult: []bool{
				true,
				false,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			s := collector.New()
			ctx := context.TODO()

			for i := range tc.pages {
				res := s.Add(ctx, tc.pages[i])
				checkIs.Equal(res, tc.expectedResult[i])
			}
		})
	}
}
