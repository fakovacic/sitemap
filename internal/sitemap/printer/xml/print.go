package printer

import (
	"context"
	"encoding/xml"
	"os"

	"github.com/fakovacic/sitemap/internal/sitemap"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Text string `xml:",chardata"`
	Loc  string `xml:"loc"`
}

func (p *printer) Print(ctx context.Context, list []sitemap.Page, outputFile string) error {
	setURLs := make([]URL, len(list))
	for i := range list {
		setURLs[i] = URL{
			Loc: list[i].URL.String(),
		}
	}

	set := URLSet{
		URLs: setURLs,
	}

	b, err := xml.MarshalIndent(set, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
