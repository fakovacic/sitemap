package sitemap

import (
	"context"
	"net/url"
)

//go:generate moq -out ./mocks/sitemap.go -pkg mocks  . Sitemap
type Sitemap interface {
	Collect(context.Context, *url.URL, int, int) error
	Output(context.Context, string) error
	Validate(context.Context, string) (*url.URL, error)
}

type sitemap struct {
	crawler   Crawler
	collector Collector
	parser    Parser
	printer   Printer
}

func New(collector Collector, crawler Crawler, parser Parser, printer Printer) Sitemap {
	return &sitemap{
		crawler:   crawler,
		collector: collector,
		parser:    parser,
		printer:   printer,
	}
}
