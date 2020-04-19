package apiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rbucchi/football-data/pkg/data/request"
)

/*
Client provides a high-level client implementation to talk to the API that football-data.org offers.

A new instance of Client will by default use the default HTTP client and no
authentication token. To configure this, Client provides methods to set the
token and the HTTP client. For more information, see the respective documentation
of SetHttpClient and SetToken, or take a look at the fluent-style companion
methods WithHttpClient and WithToken.
*/
type ApiClient struct {
	httpClient http.Client

	// Insert an API token here if you have one. It will be sent across with all requests.
	authToken string
}

func (c ApiClient) Get(out interface{}, r request.Request) error {
	var urlValues = url.Values{}
	filter, err := r.GetFilter()
	if err != nil {
		return err
	}
	if filter.Matchday != 0 {
		urlValues.Set("matchday", fmt.Sprintf("%d", filter.Matchday))
	}
	if filter.Plan != "" {
		urlValues.Set("plan", filter.Plan)
	}
	d, _, err := c.doJSON("GET", r.GetPath(), urlValues)
	if err == nil {
		err = d.Decode(&out)
	}
	return err
}

// NewClient creates a new Client instance that wraps around the given HTTP client.
//
// A call to this method is not necessary in order to create a working instance
// of Client. `new(footballdata.Client)` works just as fine.
func NewClient(h *http.Client) *ApiClient {
	return &ApiClient{httpClient: *h}
}

// SetToken sets the authentication token.
// Calling this method is *optional*.
func (c *ApiClient) SetToken(authToken string) {
	c.authToken = authToken
}

// WithToken sets the authentication token on a copy of the current Client
// instance.
//
// This method allows for easy fluent-style usage.
func (c ApiClient) WithToken(authToken string) *ApiClient {
	c.authToken = authToken
	return &c
}

// Executes an HTTP request with given parameters and on success returns the response wrapped in a JSON decoder.
func (c *ApiClient) doJSON(method string, path string, values url.Values) (j *json.Decoder, meta ResponseMeta, err error) {
	// Create request
	req := &http.Request{
		Method: method,
		URL:    resolveRelativeURL(path, values),
		Header: http.Header{},
	}

	// Set request headers
	if len(c.authToken) > 0 {
		req.Header.Set("X-Auth-Token", c.authToken)
	}
	req.Header.Set("User-Agent", "go-footballdata/0.1")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Save metadata from headers
	meta = responseMetaFromHeaders(resp.Header)

	// Expect the request to be successful, otherwise throw an error
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	// Download to buffer to allow passing back a fully prepared decoder
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	// Wrap JSON decoder around buffered data
	j = json.NewDecoder(buf)
	return
}
