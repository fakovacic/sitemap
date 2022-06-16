package parser

import (
	"context"
	"net/url"
)

func (c *parser) SetBase(ctx context.Context, base *url.URL) {
	c.base = base
}
