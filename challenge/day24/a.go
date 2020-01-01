package day24

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "24a",
	Short: "Day 24, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	w := NewWorld(challenge)

	seen := map[int]struct{}{
		w.biodiversity(): {},
	}
	for {
		w.tick()
		w.swap()
		d := w.biodiversity()

		if _, ok := seen[d]; ok {
			break
		}

		seen[d] = struct{}{}
	}

	w.dump()
	return w.biodiversity()
}
