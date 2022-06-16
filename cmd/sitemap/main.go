package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fakovacic/sitemap/internal/sitemap"
	"github.com/fakovacic/sitemap/internal/sitemap/collector"
	"github.com/fakovacic/sitemap/internal/sitemap/crawler"
	"github.com/fakovacic/sitemap/internal/sitemap/parser"
	printer "github.com/fakovacic/sitemap/internal/sitemap/printer/xml"
)

func main() {
	var (
		parallel   int
		maxDepth   int
		outputFile string
	)

	args := os.Args

	flag.IntVar(&parallel, "parallel", 4, "number of parallel workers to navigate through site")
	flag.IntVar(&maxDepth, "max-depth", 2, "max depth of url navigation recursion")
	flag.StringVar(&outputFile, "outputFile", "sitemap.xml", "output file path")
	flag.Parse()

	collector := collector.New()

	crawler := crawler.New(&http.Client{
		Timeout: 10 * time.Second,
	})

	parser := parser.New()
	printer := printer.New()

	sitemap := sitemap.New(collector, crawler, parser, printer)

	ctx := context.Background()

	// service
	baseURL, err := sitemap.Validate(ctx, args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = sitemap.Collect(ctx, baseURL, maxDepth, parallel)
	if err != nil {
		log.Fatal(err)
	}

	err = sitemap.Output(ctx, outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
