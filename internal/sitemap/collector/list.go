package collector

import (
	"context"
	"sort"

	"github.com/fakovacic/sitemap/internal/sitemap"
)

func (s *collector) List(context.Context) []sitemap.Page {
	list := make([]sitemap.Page, 0, len(s.Pages))

	for i := range s.Pages {
		if s.Pages[i].Ok {
			list = append(list, s.Pages[i])
		}
	}

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].URL.String() < list[j].URL.String()
	})

	return list
}
