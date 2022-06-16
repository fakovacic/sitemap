package collector

import (
	"context"
)

func (s *collector) AllVisited(context.Context) bool {
	s.Lock()

	all := true

	for k := range s.Pages {
		if !s.Pages[k].Visited {
			all = false
		}
	}

	s.Unlock()

	return all
}
