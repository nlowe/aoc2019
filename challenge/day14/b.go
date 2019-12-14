package day14

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "14b",
	Short: "Day 14, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	equations := map[string]equation{}
	for line := range challenge.Lines() {
		eqn := parseEquation(line)
		equations[eqn.output.name] = eqn
	}

	// not sure why this is off by one
	return sort.Search(1000000000000, func(i int) bool {
		return cost(componentFuel, i, equations, map[string]int{}) >= 1000000000000
	}) - 1
}
