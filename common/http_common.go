package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetURL make a get request on the URL
func GetURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("STATUS error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("READ BODY error: %v", err)
	}

	return data, nil
}

// ParseURL to retrieve new URL
func ParseURL(fullURL string, options ...string) string {
	dURL := ""
	u, err := url.Parse(fullURL)
	if err != nil {
		return fullURL
	}

	for _, o := range options {
		switch o {
		case "schema":
			dURL += u.Scheme + "://"
		case "host":
			dURL += u.Host
		case "path":
			dURL += u.Path
		case "query":
			dURL += "?" + u.RawQuery
		}
	}

	return dURL
}
