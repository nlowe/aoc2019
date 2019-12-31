package day23

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "23b",
	Short: "Day 23, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	inputs, outputs := bootNetwork(challenge)
	r := NewRouter(inputs, outputs)
	r.trackNat = true

	return <-r.RouteTraffic()
}
