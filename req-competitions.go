// type CompetitionsRequest struct{ Request }

// // Get retrieve list
// // func (r CompetitionsRequest) Get() (s CompetitionList, err error) {
// // 	d, err := r.Do()
// // 	if err == nil {
// // 		err = d.Decode(&s)
// // 	}
// // 	return
// // }

// // Competitions prepares a request to fetch the list of competitions for the current season or for any specified season (via the "Season" submethod).
// func (c *Client) Competitions() CompetitionsRequest {
// 	return CompetitionsRequest{c.req("competitions")}
// }

package main

import "fmt"

type CompetitionsRequest struct {
	Filter Filter
	ID     uint32
}

func (r CompetitionsRequest) path() string {
	if r.ID != 0 {
		return fmt.Sprintf("competitions/%d", r.ID)
	}
	return "competitions"
}

func (r CompetitionsRequest) filter() (f Filter, err error) {
	f = r.Filter
	return
}
