package sitemap

import (
	"context"
)

//go:generate moq -out ./mocks/collector.go -pkg mocks  . Collector
type Collector interface {
	List(context.Context) []Page
	Add(context.Context, Page) bool
	Tag(context.Context, string, bool)
	SetVisit(context.Context, string)
	AllVisited(context.Context) bool
}
