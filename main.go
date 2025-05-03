package main

import (
	"fmt"
	"sort"

	fuzzy "github.com/paul-mannino/go-fuzzywuzzy"
)

type Candidate struct {
	Text   string
	Weight float64
	Score  int
}

func main() {
	input := "fzr"
	candidates := []Candidate{
		{"fzf", 1.5, 0},
		{"fuzzy", 1.0, 0},
		{"foo", 0.5, 0},
	}

	for i := range candidates {
		candidates[i].Score = fuzzy.Ratio(input, candidates[i].Text)
		candidates[i].Score = int(float64(candidates[i].Score) * candidates[i].Weight)
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Score > candidates[j].Score
	})

	for _, c := range candidates {
		fmt.Printf("%s (score: %d)\n", c.Text, c.Score)
	}
}
