package sitemap

import "net/http"

//go:generate moq -out ./mocks/client.go -pkg mocks  . Client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}
