package printer

import (
	"github.com/fakovacic/sitemap/internal/sitemap"
)

func New() sitemap.Printer {
	return &printer{}
}

type printer struct{}
