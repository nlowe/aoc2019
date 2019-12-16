package day15

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "15b",
	Short: "Day 15, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	s, _, _ := mapShip(challenge)

	m := 0
	for tryFill(s) {
		m++
	}

	return m
}

func tryFill(s *ship) bool {
	var toFill []*tile
	s.each(func(t *tile) {
		if t.t != statusOxygenSystemFound {
			return
		}

		for _, delta := range []struct {
			x int
			y int
		}{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			if candidate := s.tileAt(t.x+delta.x, t.y+delta.y); candidate != nil && candidate.t == statusOk {
				toFill = append(toFill, candidate)
			}
		}
	})

	for _, t := range toFill {
		t.t = statusOxygenSystemFound
	}

	return len(toFill) > 0
}
