package crawler

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *crawler) Request(ctx context.Context, url *url.URL) (string, bool, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	req = req.WithContext(ctx)

	if err != nil {
		return "", false, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if !strings.Contains(contentType, "text/") {
		return "", false, errors.New("content type not valid")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", false, errors.New("status code not valid")
	}

	return string(body), true, nil
}
