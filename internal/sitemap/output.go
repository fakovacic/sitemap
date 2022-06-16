package sitemap

import (
	"context"
)

func (s *sitemap) Output(ctx context.Context, outputFile string) error {
	list := s.collector.List(ctx)

	err := s.printer.Print(ctx, list, outputFile)
	if err != nil {
		return err
	}

	return nil
}
