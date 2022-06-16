package collector

import (
	"sync"

	"github.com/fakovacic/sitemap/internal/sitemap"
)

func New() sitemap.Collector {
	return &collector{
		Pages: make(map[string]sitemap.Page),
	}
}

type collector struct {
	sync.Mutex
	Pages map[string]sitemap.Page
}
