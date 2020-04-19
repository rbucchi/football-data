package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type request struct {
	fdClient  *Client
	path      string
	urlValues url.Values
}

func (r request) Do(...filter Filter) (d *json.Decoder, err error) {
	d, _, err = r.doJSON("GET")
	r.urlValues.Set("matchday", fmt.Sprintf("%d", filter.Matchday))
	return
}

func (r request) doJSON(httpMethod string) (*json.Decoder, ResponseMeta, error) {
	return r.fdClient.doJSON(httpMethod, r.path, r.urlValues)
}
