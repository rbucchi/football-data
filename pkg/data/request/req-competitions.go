package request

import (
	"fmt"
)

type CompetitionsRequest struct {
	Filter Filter
	ID     uint32
}

func (r CompetitionsRequest) GetPath() string {
	if r.ID != 0 {
		return fmt.Sprintf("competitions/%d", r.ID)
	}
	return "competitions"
}

func (r CompetitionsRequest) GetFilter() (f Filter, err error) {
	f = r.Filter
	return
}
