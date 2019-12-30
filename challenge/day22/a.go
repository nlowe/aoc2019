package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "22a",
	Short: "Day 22, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

const (
	instrDealNewStack      = "deal into new stack"
	instrCut               = "cut "
	instrDealWithIncrement = "deal with increment "
)

func a(challenge *challenge.Input) int {
	return fastFancyShuffle(challenge, 10007, 2019)
}

func fastFancyShuffle(challenge *challenge.Input, deckSize, pos int) int {
	x := pos

	for instruction := range challenge.Lines() {
		if instruction == instrDealNewStack {
			x = fastNewStack(deckSize, x)
		} else if strings.HasPrefix(instruction, instrCut) {
			n, err := strconv.Atoi(strings.TrimPrefix(instruction, instrCut))
			if err != nil {
				panic(err)
			}

			x = fastCut(deckSize, x, n)
		} else if strings.HasPrefix(instruction, instrDealWithIncrement) {
			increment, err := strconv.Atoi(strings.TrimPrefix(instruction, instrDealWithIncrement))
			if err != nil {
				panic(err)
			}

			x = fastIncrement(deckSize, x, increment)
		}
	}

	return x
}

func fastNewStack(deckSize, x int) int {
	return safeMod(-x-1, deckSize)
}

func fastCut(deckSize, x, n int) int {
	return safeMod(x-n, deckSize)
}

func fastIncrement(deckSize, x, increment int) int {
	return safeMod(x*increment, deckSize)
}

func safeMod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}
