package day22

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "22b",
	Short: "Day 22, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

var (
	iterations = 101741582076661
)

func b(challenge *challenge.Input) int {
	return fastFancyShuffle(challenge, 119315717514047, 2020)
}
