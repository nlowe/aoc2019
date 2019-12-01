package day1

import (
	"fmt"
	"math"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "1a",
	Short: "Day 1, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(input *challenge.Input) (result int) {
	for module := range input.Lines() {
		result += FuelRequired(util.MustAtoI(module))
	}

	return
}

func FuelRequired(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}
