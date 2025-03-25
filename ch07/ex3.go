package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func (l *League) MatchResult(t1 string, s1 int, t2 string, s2 int) {
	if _, ok := l.Teams[t1]; !ok {
		e := fmt.Errorf("%s not found", t1)
		fmt.Println(e.Error())
		return
	}
	if _, ok := l.Teams[t2]; !ok {
		e := fmt.Errorf("%s not found", t2)
		fmt.Println(e.Error())
		return
	}
	if s1 == s2 {
		return
	}
	if s1 > s2 {
		l.Wins[t1]++
	} else {
		l.Wins[t2]++
	}
}

// Ranking returns a slice of the team names in order of wins
func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	fmt.Println(results)
	for _, result := range results {
		_, _ = io.WriteString(w, result)
		_, _ = w.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Teams: map[string]Team{
			"Team 1": {
				Name:    "Team 1",
				Players: []string{"Player 1", "Player 2", "Player 3"},
			},
			"Team 2": {
				Name:    "Team 2",
				Players: []string{"Player 4", "Player 5", "Player 6"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Team 1", 1, "Team 2", 2)
	l.MatchResult("Team 1", 2, "Team 2", 1)
	l.MatchResult("Team 1", 3, "Team 2", 10)
	RankPrinter(&l, os.Stdout)
	RankPrinter()
}
