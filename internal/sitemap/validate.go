package sitemap

import (
	"context"
	"errors"
	"net/url"
)

func (s *sitemap) Validate(ctx context.Context, link string) (*url.URL, error) {
	baseURL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	if baseURL.Scheme == "" {
		return nil, errors.New("scheme is empty")
	}

	if baseURL.Host == "" {
		return nil, errors.New("host is empty")
	}

	return baseURL, nil
}
