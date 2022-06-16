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

func TestOutput(t *testing.T) {
	cases := []struct {
		it         string
		outputFile string

		collectorResult []sitemap.Page

		printerInputPages   []sitemap.Page
		printerInputOutFile string
		printerError        error

		expectedError string
	}{
		{
			it:         "it generate links to file",
			outputFile: "mock-file",
			collectorResult: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			printerInputPages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			printerInputOutFile: "mock-file",
		},
		{
			it:         "it return error on generate ",
			outputFile: "mock-file",
			collectorResult: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			printerInputPages: []sitemap.Page{
				{
					URL: &url.URL{
						Scheme: "https",
						Host:   "mock.com",
					},
				},
			},
			printerInputOutFile: "mock-file",
			printerError:        errors.New("mock-error"),
			expectedError:       "mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			collector := &mocks.CollectorMock{
				ListFunc: func(ctx context.Context) []sitemap.Page {
					return tc.collectorResult
				},
			}

			printer := &mocks.PrinterMock{
				PrintFunc: func(ctx context.Context, pages []sitemap.Page, outFile string) error {
					checkIs.Equal(pages, tc.printerInputPages)
					checkIs.Equal(outFile, tc.printerInputOutFile)

					return tc.printerError
				},
			}

			s := sitemap.New(collector, nil, nil, printer)
			ctx := context.TODO()

			err := s.Output(ctx, tc.outputFile)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
