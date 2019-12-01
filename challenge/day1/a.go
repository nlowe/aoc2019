package day1

import (
    "fmt"
    "github.com/nlowe/aoc2019/challenge"
    "github.com/spf13/cobra"
)

var A = &cobra.Command{
    Use: "1a",
    Short: "Day 1, Problem A",
    Run: func(_ *cobra.Command, _ []string) {
        fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
    },
}

func a(input *challenge.Input) int {
    return 0
}
