package main

import "fmt"

type Request interface {
	path() string
	filter() (Filter, error)
}

type Filter struct {
	Matchday     uint32
	Season       string
	Status       StatusType
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
