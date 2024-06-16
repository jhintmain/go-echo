package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name    string
	Age     int
	Score   int
	Percent float64
}

type Players []Player

func (p Players) Len() int {
	return len(p)
}
func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Players) Less(i, j int) bool {
	return p[i].Score < p[j].Score
}

func main() {
	users := Players{
		{"니통키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}

	sort.Sort(users)
	fmt.Print(users)
}
