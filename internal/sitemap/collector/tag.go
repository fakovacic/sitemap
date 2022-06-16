package collector

import (
	"context"
)

func (s *collector) Tag(ctx context.Context, pageURL string, ok bool) {
	s.Lock()

	val, found := s.Pages[pageURL]
	if found {
		val.Ok = ok
		s.Pages[pageURL] = val
	}

	s.Unlock()
}
