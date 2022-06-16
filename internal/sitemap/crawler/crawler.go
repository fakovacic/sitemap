package crawler

import (
	"github.com/fakovacic/sitemap/internal/sitemap"
)

func New(client sitemap.Client) sitemap.Crawler {
	return &crawler{
		client: client,
	}
}

type crawler struct {
	client sitemap.Client
}
