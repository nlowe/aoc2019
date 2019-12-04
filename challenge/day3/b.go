package day3

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "3b",
	Short: "Day 3, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(input *challenge.Input) int {
	wires := input.Lines()
	w1 := makeSegments(<-wires)
	w2 := makeSegments(<-wires)

	w1d := 0
	for _, w1s := range w1 {
		w2d := 0
		for _, w2s := range w2 {
			x, y, intersects := w1s.intersection(w2s)
			if intersects && ManhattanDistance(0, 0, x, y) > 0 {
				return w1d + w1s.partialLength(x, y) + w2d + w2s.partialLength(x, y)
			}

			w2d += w2s.length()
		}

		w1d += w1s.length()
	}

	panic("no solution found")
}
