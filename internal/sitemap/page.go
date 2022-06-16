package sitemap

import "net/url"

type Page struct {
	URL     *url.URL
	Visited bool
	Ok      bool
}
