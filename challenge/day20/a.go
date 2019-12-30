package day20

import (
	"fmt"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "20a",
	Short: "Day 20, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	d := NewDonut(challenge)

	_, dist, found := astar.Path(d.start, d.end)
	if !found {
		panic("no solution")
	}

	return int(dist)
}
