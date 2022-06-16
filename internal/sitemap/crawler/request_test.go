package crawler_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/fakovacic/sitemap/internal/sitemap/crawler"
	"github.com/fakovacic/sitemap/internal/sitemap/mocks"
	"github.com/matryer/is"
)

func TestList(t *testing.T) {
	cases := []struct {
		it  string
		url *url.URL

		clientDoResponse *http.Response
		clientDoError    error

		expectedResultOk   bool
		expectedResultBody string
		expectedError      string
	}{
		{
			it: "it return body response",
			url: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			clientDoResponse: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`<body>mock</body>`)),
				Header: map[string][]string{
					"Content-Type": {
						"text/",
					},
				},
			},
			expectedResultOk:   true,
			expectedResultBody: "<body>mock</body>",
		},
		{
			it: "it return body response",
			url: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			clientDoResponse: &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`<body>mock</body>`)),
				Header: map[string][]string{
					"Content-Type": {
						"text/",
					},
				},
			},
			expectedError: "status code not valid",
		},
		{
			it: "it return body response",
			url: &url.URL{
				Scheme: "https",
				Host:   "mock.com",
			},
			clientDoResponse: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`<body>mock</body>`)),
				Header: map[string][]string{
					"Content-Type": {
						"image/",
					},
				},
			},
			expectedError: "content type not valid",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			mockClient := &mocks.ClientMock{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return tc.clientDoResponse, tc.clientDoError
				},
			}

			s := crawler.New(mockClient)
			ctx := context.TODO()

			body, ok, err := s.Request(ctx, tc.url)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}

			checkIs.Equal(ok, tc.expectedResultOk)
			checkIs.Equal(body, tc.expectedResultBody)

		})
	}
}
