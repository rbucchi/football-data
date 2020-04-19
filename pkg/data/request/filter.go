package request

import "fmt"

type Filter struct {
	Matchday     uint32
	Season       string
	Status       string // StatusTypes
	Venue        string // Venue
	DateFrom     string
	DateTo       string
	Stage        string
	Plan         string
	Competitions string
	Group        string
	Limit        uint32
	StandingType string
}

type FilterError struct {
	Msg string
}

func (r *FilterError) Error() string {
	return fmt.Sprintf("Msg %v", r.Msg)
}
