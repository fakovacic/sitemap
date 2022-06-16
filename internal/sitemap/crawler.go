package sitemap

import (
	"context"
	"net/url"
)

//go:generate moq -out ./mocks/crawler.go -pkg mocks  . Crawler
type Crawler interface {
	Request(context.Context, *url.URL) (string, bool, error)
}
