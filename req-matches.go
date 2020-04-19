package main

type MatchesRequest struct{ request }

func (c *Client) Matches(competitionId uint32) MatchesRequest {
	r := MatchesRequest{c.req("competitions/%d/matches", competitionId)}
	return r
}

// Get retrieve list
func (r MatchesRequest) Get() (s MatchList, err error) {
	d, err := r.Do()

	if err == nil {
		err = d.Decode(&s)
	}
	return
}
