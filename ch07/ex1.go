package main

import "fmt"

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(t1 string, s1 int, t2 string, s2 int) {
	if _, ok := l.Teams[t1]; !ok {
		_ = fmt.Sprintf("Team %s not found", t1)
		return
	}
	if _, ok := l.Teams[t2]; !ok {
		_ = fmt.Sprintf("Team %s not found", t2)
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

func (l *League) Ranking() []string {
	//logic
}

func main() {}
