package sitemap

import (
	"context"
	"fmt"
	"net/url"
	"sync"
)

type fetchUrl struct {
	URL  *url.URL
	Dept int
}

func (s *sitemap) Collect(ctx context.Context, baseURL *url.URL, dept, workers int) error {
	s.parser.SetBase(ctx, baseURL)

	fetch := make(chan fetchUrl, workers)
	visitPages := make(chan fetchUrl)

	_ = s.collector.Add(ctx, Page{
		URL: baseURL,
	})

	fetch <- fetchUrl{
		URL:  baseURL,
		Dept: dept,
	}

	var w sync.WaitGroup

	w.Add(1)

	go func() {
		for {
			select {
			case p := <-fetch:
				w.Add(1)

				go func(p fetchUrl) {
					defer w.Done()

					pages, err := s.collect(ctx, p)
					if err != nil {
						fmt.Println(fmt.Sprintf("%s : %v", p.URL.String(), err))

						visitPages <- p

						return
					}

					for i := range pages {
						newLink := s.collector.Add(ctx, pages[i])
						if newLink {
							fetch <- fetchUrl{
								URL:  pages[i].URL,
								Dept: dept - 1,
							}
						}
					}

					visitPages <- p
				}(p)

			case p := <-visitPages:
				s.collector.SetVisit(ctx, p.URL.String())

				if s.collector.AllVisited(ctx) {
					close(fetch)
					w.Done()

					return
				}
			}
		}
	}()

	w.Wait()

	return nil
}

func (s *sitemap) collect(ctx context.Context, p fetchUrl) ([]Page, error) {
	content, ok, err := s.crawler.Request(ctx, p.URL)

	s.collector.Tag(ctx, p.URL.String(), ok)

	if err != nil {
		return nil, err
	}

	if p.Dept == 0 {
		return nil, nil
	}

	pages, err := s.parser.Parse(ctx, content)
	if err != nil {
		return nil, err
	}

	return pages, err
}
