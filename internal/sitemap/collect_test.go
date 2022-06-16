package sitemap_test

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/fakovacic/sitemap/internal/sitemap/mocks"
	"github.com/matryer/is"
)

func TestCollectSingle(t *testing.T) {
	cases := []struct {
		it string

		url     *url.URL
		dept    int
		workers int

		// crawler
		crawlerInput    *url.URL
		crawlerResult   string
		crawlerResultOk bool
		crawlerError    error

		// parser
		parserInputSetBase *url.URL

		parserParseInput  string
		parserParseResult []sitemap.Page
		parserParseError  error

		// collector
		collectorAddInput  sitemap.Page
		collectorAddResult bool

		collectorTagInputPage string
		collectorTagInputOk   bool

		collectorSetVisitInput string

		collectorAllVisitedResult bool

		expectedError string
	}{
		{
			it: "it visit one link",
			url: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			dept:    1,
			workers: 1,

			parserInputSetBase: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},

			collectorAddInput: sitemap.Page{
				URL: &url.URL{
					Scheme: "https",
					Host:   "mock.com",
				},
			},
			crawlerInput: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			crawlerResult:   "mock-body",
			crawlerResultOk: true,

			collectorTagInputPage: "https://mock.com",
			collectorTagInputOk:   true,

			parserParseInput:  "mock-body",
			parserParseResult: []sitemap.Page{},

			collectorSetVisitInput: "https://mock.com",

			collectorAllVisitedResult: true,
		},
		{
			it: "it return error on parse",
			url: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			dept:    1,
			workers: 1,

			parserInputSetBase: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},

			collectorAddInput: sitemap.Page{
				URL: &url.URL{
					Scheme: "https",
					Host:   "mock.com",
				},
			},
			crawlerInput: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			crawlerResult:   "mock-body",
			crawlerResultOk: true,

			collectorTagInputPage: "https://mock.com",
			collectorTagInputOk:   true,

			parserParseInput: "mock-body",
			parserParseError: errors.New("mock-error"),

			collectorSetVisitInput: "https://mock.com",

			collectorAllVisitedResult: true,
		},
		{
			it: "it return error on parse",
			url: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			dept:    1,
			workers: 1,

			parserInputSetBase: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},

			collectorAddInput: sitemap.Page{
				URL: &url.URL{
					Scheme: "https",
					Host:   "mock.com",
				},
			},
			crawlerInput: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			crawlerResult:   "mock-body",
			crawlerResultOk: true,
			crawlerError:    errors.New("mock-error"),

			collectorTagInputPage: "https://mock.com",
			collectorTagInputOk:   true,

			collectorSetVisitInput: "https://mock.com",

			collectorAllVisitedResult: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			collector := &mocks.CollectorMock{
				AddFunc: func(ctx context.Context, page sitemap.Page) bool {
					checkIs.Equal(page, tc.collectorAddInput)

					return tc.collectorAddResult
				},
				TagFunc: func(ctx context.Context, page string, ok bool) {
					checkIs.Equal(page, tc.collectorTagInputPage)
					checkIs.Equal(ok, tc.collectorTagInputOk)
				},
				SetVisitFunc: func(ctx context.Context, page string) {
					checkIs.Equal(page, tc.collectorSetVisitInput)
				},
				AllVisitedFunc: func(ctx context.Context) bool {
					return tc.collectorAllVisitedResult
				},
			}

			crawler := &mocks.CrawlerMock{
				RequestFunc: func(ctx context.Context, url *url.URL) (string, bool, error) {
					checkIs.Equal(url, tc.crawlerInput)

					return tc.crawlerResult, tc.crawlerResultOk, tc.crawlerError
				},
			}

			parser := &mocks.ParserMock{
				SetBaseFunc: func(ctx context.Context, link *url.URL) {
					checkIs.Equal(link, tc.parserInputSetBase)
				},
				ParseFunc: func(ctx context.Context, link string) ([]sitemap.Page, error) {
					checkIs.Equal(link, tc.parserParseInput)

					return tc.parserParseResult, tc.parserParseError
				},
			}

			s := sitemap.New(collector, crawler, parser, nil)
			ctx := context.TODO()

			err := s.Collect(ctx, tc.url, tc.dept, tc.workers)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}

func TestCollectMultiple(t *testing.T) {

}
