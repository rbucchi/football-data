package main

type CompetitionsRequest struct{ request }

// Get retrieve list
func (r CompetitionsRequest) Get() (s CompetitionList, err error) {
	d, err := r.Do()
	if err == nil {
		err = d.Decode(&s)
	}
	return
}

// Competitions prepares a request to fetch the list of competitions for the current season or for any specified season (via the "Season" submethod).
func (c *Client) Competitions() CompetitionsRequest {
	return CompetitionsRequest{c.req("competitions")}
}
