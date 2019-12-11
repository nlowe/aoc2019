package day10

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const targetIndex = 200

var B = &cobra.Command{
	Use:   "10b",
	Short: "Day 10, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	asteroids := makeMap(challenge)

	station, _, targets := findStation(asteroids)

	sort.Slice(targets, func(i, j int) bool {
		aTheta := station.angleTo(targets[i])
		bTheta := station.angleTo(targets[j])

		if aTheta == bTheta {
			return station.distanceTo(targets[i]) < station.distanceTo(targets[j])
		}

		return aTheta < bTheta
	})

	return int(targets[targetIndex-1].x)*100 + int(targets[targetIndex-1].y)
}
