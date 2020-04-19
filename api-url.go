package main

import "net/url"

const (
	baseURL = "http://api.football-data.org/v2/"
)

func resolveRelativeURL(path string, values url.Values) *url.URL {
	if values == nil {
		values = url.Values{}
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	ru := &url.URL{Path: path, RawQuery: values.Encode()}

	return u.ResolveReference(ru)
}
