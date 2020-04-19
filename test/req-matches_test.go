package test

import (
	"errors"
	"net/http"
	"testing"

	apiclient "github.com/rbucchi/football-data/pkg/api-client"
	"github.com/rbucchi/football-data/pkg/data/business"
	"github.com/rbucchi/football-data/pkg/data/request"
	"github.com/stretchr/testify/assert"
)

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestCreateRepoSuccess(t *testing.T) {
	client := new(apiclient.ApiClient)
	client.HttpClient = MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error from web server",
		)
	}
	var matches business.MatchList
	err := client.Get(
		&matches,
		request.MatchesRequest{
			Filter:      request.Filter{Matchday: 1, Plan: ""},
			Competition: 2019,
		},
	)
	assert.NotNil(t, err)
	// assert.Nil(t, matches)
	// assert.EqualValues(t, "Test Name", resp.Name)
}
