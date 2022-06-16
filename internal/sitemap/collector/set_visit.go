package collector

import (
	"context"
)

func (s *collector) SetVisit(ctx context.Context, pageURL string) {
	s.Lock()

	val, ok := s.Pages[pageURL]
	if ok {
		val.Visited = true
		s.Pages[pageURL] = val
	}

	s.Unlock()
}
