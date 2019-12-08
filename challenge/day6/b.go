package day6

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const (
	ObjectYou   = "YOU"
	ObjectSanta = "SAN"
)

var B = &cobra.Command{
	Use:   "6b",
	Short: "Day 6, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	bodies := makeMap(challenge)

	for you, yt := bodies[ObjectYou], 0; you != nil; you, yt = you.Orbits, yt+1 {
		for santa, st := bodies[ObjectSanta], 0; santa != nil; santa, st = santa.Orbits, st+1 {
			if santa == you {
				// Don't include the orbital bodies you started around
				return yt + st - 2
			}
		}
	}

	panic("no solution")
}
