package sitemap_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/matryer/is"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		it   string
		link string

		expectedURL   *url.URL
		expectedError string
	}{
		{
			it:   "it returns url",
			link: "https://mock.com",
			expectedURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			s := sitemap.New(nil, nil, nil, nil)

			url, err := s.Validate(context.TODO(), tc.link)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}

			checkIs.Equal(url, tc.expectedURL)
		})
	}
}
