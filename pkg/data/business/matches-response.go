package business

import "time"

type MatchesResponse struct {
	Count   int `json:"count"`
	Filters struct {
		Matchday string `json:"matchday"`
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Matches []struct {
		ID     int `json:"id"`
		Season struct {
			ID              int    `json:"id"`
			StartDate       string `json:"startDate"`
			EndDate         string `json:"endDate"`
			CurrentMatchday int    `json:"currentMatchday"`
		} `json:"season"`
		UtcDate     time.Time `json:"utcDate"`
		Status      string    `json:"status"`
		Matchday    int       `json:"matchday"`
		Stage       string    `json:"stage"`
		Group       string    `json:"group"`
		LastUpdated time.Time `json:"lastUpdated"`
		Odds        struct {
			Msg string `json:"msg"`
		} `json:"odds"`
		Score struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"fullTime"`
			HalfTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"halfTime"`
			ExtraTime struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"extraTime"`
			Penalties struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"penalties"`
		} `json:"score"`
		HomeTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"awayTeam"`
		Referees []struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Nationality interface{} `json:"nationality"`
		} `json:"referees"`
	} `json:"matches"`
}
