package main

type CompetitionRequest struct{ request }

// Get retrieve list
func (r CompetitionRequest) Get() (s Competition, err error) {
	d, err := r.Do()
	if err == nil {
		err = d.Decode(&s)
	}
	return
}

// Competition get
func (c *Client) Competition(competitionID uint32) CompetitionRequest {
	return CompetitionRequest{c.req("competitions/%d", competitionID)}
}
