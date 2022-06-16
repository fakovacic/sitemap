package collector_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/fakovacic/sitemap/internal/sitemap/collector"
	"github.com/matryer/is"
)

func TestAllVisited(t *testing.T) {
	cases := []struct {
		it    string
		pages []sitemap.Page

		expectedResult bool
	}{
		{
			it: "it check all visited pages in collector - true",
			pages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
					Visited: true,
				},
			},
			expectedResult: true,
		},
		{
			it: "it check all visited pages in collector - false",
			pages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			expectedResult: false,
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

			allVisited := s.AllVisited(ctx)

			checkIs.Equal(allVisited, tc.expectedResult)
		})
	}
}
