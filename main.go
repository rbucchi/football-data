package main

import (
	"fmt"
	"sort"
)

func main() {
	client := new(Client).
		WithToken("81df68ad98c746d6a1daebbddab8c23b")
	// printCompetitions(client)
	matches, err := client.Matches(2019).SetFilter(Filter{Matchday: 1}).Get()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, m := range matches.Matches {
		fmt.Println(m)
	}

}

func printCompetitions(client *Client) {
	// Get list of seasons...
	competitions, err := client.Competitions().Get()
	if err != nil {
		panic(err)
	}

	sortedC := competitions.Competitions
	sort.Slice(sortedC, func(i, j int) bool {
		return sortedC[i].Area.Name < sortedC[j].Area.Name
	})

	// ...and print them
	for _, competition := range sortedC {
		fmt.Println(competition)
		// fmt.Println(competition.ID, competition.Area.Name, competition.Name, competition.Plan)
	}
}
