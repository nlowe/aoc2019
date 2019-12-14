package day14

import (
	"fmt"

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

	lo := 1000000000000 / cost(componentFuel, 1, equations, map[string]int{})
	hi := 1000000000000

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if cost(componentFuel, mid, equations, map[string]int{}) <= 1000000000000 {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}
