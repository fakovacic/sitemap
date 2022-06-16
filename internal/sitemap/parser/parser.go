package parser

import (
	"net/url"

	"github.com/fakovacic/sitemap/internal/sitemap"
)

func New() sitemap.Parser {
	return &parser{}
}

type parser struct {
	base *url.URL
}
