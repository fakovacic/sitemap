package collector

import (
	"context"

	"github.com/fakovacic/sitemap/internal/sitemap"
)

func (s *collector) Add(ctx context.Context, page sitemap.Page) bool {
	s.Lock()

	newPage := false

	_, ok := s.Pages[page.URL.String()]
	if !ok {
		s.Pages[page.URL.String()] = page

		newPage = true
	}

	s.Unlock()

	return newPage
}
