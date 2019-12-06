package day6

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const orbitSymbol = ")"

type body struct {
	Name   string
	Orbits *body
}

var A = &cobra.Command{
	Use:   "6a",
	Short: "Day 6, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	bodies := makeMap(challenge)

	orbits := 0
	for _, b := range bodies {
		orbits += countOrbits(b)
	}

	return orbits
}

func countOrbits(b *body) int {
	if b.Orbits == nil {
		return 0
	}

	return 1 + countOrbits(b.Orbits)
}

func makeMap(challenge *challenge.Input) map[string]*body {
	bodies := map[string]*body{}

	for orbit := range challenge.Lines() {
		parts := strings.Split(orbit, orbitSymbol)

		obj, ok := bodies[parts[1]]
		if !ok {
			obj = &body{Name: parts[1]}
			bodies[obj.Name] = obj
		}

		around, ok := bodies[parts[0]]
		if !ok {
			around = &body{Name: parts[0]}
			bodies[around.Name] = around
		}

		obj.Orbits = around
	}

	return bodies
}
