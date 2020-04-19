package main

import "fmt"

type MatchesRequest struct {
	Filter      Filter
	Competition uint32
}

func (mr MatchesRequest) path() string {
	if mr.Competition != 0 {
		return fmt.Sprintf("competitions/%d/matches", mr.Competition)
	}
	return "matches"
}

func (mr MatchesRequest) filter() (f Filter, err error) {
	f = mr.Filter
	if mr.Filter.Plan != "" {
		err = &FilterError{Msg: "cannot filter matches by plan"}
	}
	return
}
