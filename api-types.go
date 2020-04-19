package main

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
