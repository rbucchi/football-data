package main

import (
	"fmt"

	apiclient "github.com/rbucchi/football-data/pkg/api-client"
	"github.com/rbucchi/football-data/pkg/data/business"
	"github.com/rbucchi/football-data/pkg/data/request"
)

func main() {
	client := new(apiclient.ApiClient).WithToken("<TOKEN-PLACEHOLDER>")
	var matches business.MatchesResponse
	err := client.Get(
		&matches,
		request.MatchesRequest{
			Filter:      request.Filter{Matchday: 1},
			Competition: 2019,
		},
	)
	if err == nil {
		fmt.Println(matches)
	} else {
		fmt.Println(err)
		panic(err)
	}
}
