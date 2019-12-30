package day20

import (
	"fmt"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "20b",
	Short: "Day 20, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	d := NewDonut(challenge)
	d.gatesChangeLevels = true

	_, dist, found := astar.Path(d.start, d.end)
	if !found {
		panic("no solution")
	}

	return int(dist)
}
