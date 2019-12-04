package day4

import (
	"fmt"
	"strconv"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "4b",
	Short: "Day 4, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(input *challenge.Input) (count int) {
	low, high := parseRange(input)

	for pw := low; pw <= high; pw++ {
		if isValidPassword(pw) && repeatRule(pw) {
			count++
		}
	}

	return
}

func repeatRule(pw int) bool {
	s := strconv.Itoa(pw)

	target := s[0]
	chain := 1
	matching := false
	for i := 1; i < len(s); i++ {
		if s[i] == target {
			chain++
			matching = chain == 2
		} else if chain == 2 {
			return true
		} else {
			target = s[i]
			chain = 1
		}
	}

	return matching
}
