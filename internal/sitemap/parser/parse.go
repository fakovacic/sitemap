package parser

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/fakovacic/sitemap/internal/sitemap"
)

func (c *parser) Parse(ctx context.Context, content string) ([]sitemap.Page, error) {
	rgxHref, err := regexp.Compile(`<a(.*)href=\"([^\\"]*)\"`)
	if err != nil {
		return nil, err
	}

	rgxUri, err := regexp.Compile(`[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	if err != nil {
		return nil, err
	}

	links := rgxHref.FindAllString(content, -1)
	if links == nil {
		return nil, errors.New("not found links")
	}

	list := make([]sitemap.Page, 0, len(links))

	for i := range links {
		linkURL := rgxUri.FindString(links[i])

		if linkURL == "" {
			continue
		}

		validURL, ok := c.clearLink(ctx, linkURL)
		if ok {
			list = append(list, sitemap.Page{
				URL: validURL,
			})
		}
	}

	return list, nil
}

func (c *parser) clearLink(ctx context.Context, linkURL string) (*url.URL, bool) {
	var err error

	linkURL = strings.TrimSuffix(linkURL, "/")

	linkURL, err = url.QueryUnescape(linkURL)
	if err != nil {
		return nil, false
	}

	if strings.Contains(linkURL, "tel:") {
		return nil, false
	}

	if strings.Contains(linkURL, "mailto:") {
		return nil, false
	}

	if strings.Contains(linkURL, "#") {
		return nil, false
	}

	u, err := url.Parse(linkURL)
	if err != nil {
		return nil, false
	}

	if u.Host != "" && u.Host != c.base.Host {
		return nil, false
	}

	if u.Host == "" {
		u.Host = c.base.Host
	}

	if u.Scheme == "" {
		u.Scheme = c.base.Scheme
	}

	u.RawQuery = ""

	ok := c.validatePath(ctx, u.Path)
	if !ok {
		return nil, false
	}

	validURL, err := url.ParseRequestURI(u.String())
	if err != nil {
		return nil, false
	}

	return validURL, true
}

func (c *parser) validatePath(ctx context.Context, path string) bool {
	pathParts := strings.Split(path, "/")

	match, _ := regexp.MatchString(`^[a-zA-Z0-9](?:[a-zA-Z0-9 ._-]*[a-zA-Z0-9])?\.[a-zA-Z0-9_-]+$`, pathParts[len(pathParts)-1])

	return !match
}
