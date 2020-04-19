package main

import "fmt"

func main() {
	client := new(Client).
		WithToken("81df68ad98c746d6a1daebbddab8c23b")
	var matches MatchList
	err := client.Get(
		&matches,
		MatchesRequest{
			Filter:      Filter{Matchday: 1, Plan: ""},
			Competition: 2019,
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, m := range matches.Matches {
		fmt.Println(m)
	}
	var competition Competition
	err = client.Get(
		&competition,
		CompetitionsRequest{
			// Filter: Filter{Matchday: 1, Plan: "sdc"},
			ID: 2019,
		},
	)
	fmt.Println()
	fmt.Println(competition)

	var competitions CompetitionList
	err = client.Get(
		&competitions,
		CompetitionsRequest{
		// Filter: Filter{Matchday: 1, Plan: "sdc"},
		},
	)
	fmt.Println()
	fmt.Println(competitions)
}

// func printCompetitions(client *Client) {
// 	// Get list of seasons...
// 	competitions, err := client.Competitions().Get()
// 	if err != nil {
// 		panic(err)
// 	}

// 	sortedC := competitions.Competitions
// 	sort.Slice(sortedC, func(i, j int) bool {
// 		return sortedC[i].Area.Name < sortedC[j].Area.Name
// 	})

// 	// ...and print them
// 	for _, competition := range sortedC {
// 		fmt.Println(competition)
// 		// fmt.Println(competition.ID, competition.Area.Name, competition.Name, competition.Plan)
// 	}
// }
