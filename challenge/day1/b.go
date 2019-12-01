package day1

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "1b",
	Short: "Day 1, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(input *challenge.Input) (result int) {
	for module := range input.Lines() {
		moduleCost := FuelRequired(util.MustAtoI(module))

		result += moduleCost + fuelCost(moduleCost)
	}

	return
}

func fuelCost(f int) int {
	cost := FuelRequired(f)
	if cost <= 0 {
		return 0
	}

	return cost + fuelCost(cost)

}
