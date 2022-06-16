package sitemap

import (
	"context"
	"net/url"
)

//go:generate moq -out ./mocks/parser.go -pkg mocks  . Parser
type Parser interface {
	SetBase(context.Context, *url.URL)
	Parse(context.Context, string) ([]Page, error)
}
