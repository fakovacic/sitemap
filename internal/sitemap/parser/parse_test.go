package parser_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/fakovacic/sitemap/internal/sitemap/parser"
	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	cases := []struct {
		it      string
		content string
		baseURL *url.URL

		expectedResult []sitemap.Page
		expectedError  string
	}{
		{
			it: "it return parsed pages",
			baseURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			content: `<body>
				<a href="https://mock.com">Mock</a>
			</body>`,
			expectedResult: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
		},
		{
			it: "# link",
			baseURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			content: `<body>
				<a href="#">Mock</a>
			</body>`,
			expectedResult: []sitemap.Page{},
		},
		{
			it: "# link",
			baseURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			content: `<body>
				<a href="https://mock.com#test">Mock</a>
			</body>`,
			expectedResult: []sitemap.Page{},
		},
		{
			it: "tel link",
			baseURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			content: `<body>
				<a href="tel:123456">Mock</a>
			</body>`,
			expectedResult: []sitemap.Page{},
		},
		{
			it: "mailto link",
			baseURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			content: `<body>
				<a href="mailto:test@test.com">Mock</a>
			</body>`,
			expectedResult: []sitemap.Page{},
		},
		{
			it: "no links",
			baseURL: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			content:       `<body></body>`,
			expectedError: "not found links",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			s := parser.New()
			ctx := context.TODO()

			s.SetBase(ctx, tc.baseURL)
			pages, err := s.Parse(ctx, tc.content)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(pages, tc.expectedResult)
		})
	}
}
