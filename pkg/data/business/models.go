package business

import (
	"time"

	"github.com/rbucchi/football-data/pkg/data/request"
)

// Area type
type Area struct {
	ID           uint32
	Name         string
	CountryCode  string
	ParentAreaID uint32
	ParentArea   string
	ChildAreas   []Area
}

type AreaReq struct {
	Count   uint32
	Filters []request.Filter
	Area    []Area
}
type StatusType string

const (
	Scheduled StatusType = "SCHEDULED"
	Live      StatusType = "LIVE"
	InPlay    StatusType = "IN_PLAY"
	Paused    StatusType = "PAUSED"
	Finished  StatusType = "FINISHED"
	Postponed StatusType = "POSTPONED"
	Suspended StatusType = "SUSPENDED"
	Canceled  StatusType = "CANCELED"
)

// Contains the list of competitions returned by the API.
type CompetitionList struct {
	Filters      request.Filter
	Competitions []Competition
}
type MatchList struct {
	Filters interface{}
	Matches []Match
	Count   uint32
}

// Contains information about a competition.
type Competition struct {
	Id                       uint64
	Name                     string
	Code                     string
	Plan                     string
	EmblemURL                string
	NumberOfAvailableSeasons uint16
	Area                     Area
	LastUpdated              time.Time
}
type Season struct {
	ID              int      `json:"id"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CurrentMatchday int      `json:"currentMatchday"`
	AvailableStages []string `json:"availableStages"`
}
type Coach struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CountryOfBirth string `json:"countryOfBirth"`
	Nationality    string `json:"nationality"`
}
type Captain struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ShirtNumber int    `json:"shirtNumber"`
}
type LineupPlayer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	ShirtNumber int    `json:"shirtNumber"`
}
type MatchTeam struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Coach   Coach          `json:"coach"`
	Captain Captain        `json:"captain"`
	Lineup  []LineupPlayer `json:"lineup"`
	Bench   []LineupPlayer `json:"bench"`
}
type PartialScore struct {
	HomeTeam int `json:"homeTeam"`
	AwayTeam int `json:"awayTeam"`
}
type Score struct {
	Winner    string       `json:"winner"`
	Duration  string       `json:"duration"`
	FullTime  PartialScore `json:"fullTime"`
	HalfTime  PartialScore `json:"halfTime"`
	ExtraTime PartialScore `json:"extraTime"`
	Penalties PartialScore `json:"penalties"`
}
type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Goals struct {
	Minute    int         `json:"minute"`
	ExtraTime interface{} `json:"extraTime"`
	Type      string      `json:"type"`
	Team      MatchTeam   `json:"team"`
	Scorer    Player      `json:"scorer"`
	Assist    Player      `json:"assist"`
}
type Bookings struct {
	Minute int       `json:"minute"`
	Team   MatchTeam `json:"team"`
	Player Player    `json:"player"`
	Card   string    `json:"card"`
}
type Substitutions struct {
	Minute    int       `json:"minute"`
	Team      MatchTeam `json:"team"`
	PlayerOut Player    `json:"playerOut"`
	PlayerIn  Player    `json:"playerIn"`
}
type Referees struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Nationality interface{} `json:"nationality"`
}
type Match struct {
	ID            int             `json:"id"`
	Competition   Competition     `json:"competition"`
	Season        Season          `json:"season"`
	UtcDate       time.Time       `json:"utcDate"`
	Status        string          `json:"status"`
	Minute        interface{}     `json:"minute"`
	Attendance    int             `json:"attendance"`
	Venue         string          `json:"venue"`
	Matchday      int             `json:"matchday"`
	Stage         string          `json:"stage"`
	Group         string          `json:"group"`
	LastUpdated   time.Time       `json:"lastUpdated"`
	HomeTeam      MatchTeam       `json:"homeTeam"`
	AwayTeam      MatchTeam       `json:"awayTeam"`
	Score         Score           `json:"score"`
	Goals         []Goals         `json:"goals"`
	Bookings      []Bookings      `json:"bookings"`
	Substitutions []Substitutions `json:"substitutions"`
	Referees      []Referees      `json:"referees"`
}
