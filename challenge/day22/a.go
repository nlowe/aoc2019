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

var deckSize = 10007

func a(challenge *challenge.Input) int {
	deck := fancyShuffle(challenge)

	for i, v := range deck {
		if v == 2019 {
			return i
		}
	}

	panic("no solution")
}

func fancyShuffle(challenge *challenge.Input) []int {
	deck := make([]int, deckSize)
	for i := range deck {
		deck[i] = i
	}

	for instruction := range challenge.Lines() {
		if instruction == instrDealNewStack {
			deck = dealIntoNewStack(deck)
		} else if strings.HasPrefix(instruction, instrCut) {
			n, err := strconv.Atoi(strings.TrimPrefix(instruction, instrCut))
			if err != nil {
				panic(err)
			}

			deck = cut(deck, n)
		} else if strings.HasPrefix(instruction, instrDealWithIncrement) {
			increment, err := strconv.Atoi(strings.TrimPrefix(instruction, instrDealWithIncrement))
			if err != nil {
				panic(err)
			}

			deck = dealWithIncrement(deck, increment)
		}
	}

	return deck
}

func dealIntoNewStack(deck []int) []int {
	newStack := make([]int, len(deck))
	for i, v := range deck {
		newStack[len(newStack)-i-1] = v
	}

	return newStack
}

func cut(deck []int, n int) []int {
	newStack := make([]int, len(deck))
	for i := range deck {
		if n > 0 {
			newStack[i] = deck[(n+i)%len(deck)]
		} else {
			newStack[i] = deck[(len(deck)+i+n)%len(deck)]
		}
	}

	return newStack
}

func dealWithIncrement(deck []int, increment int) []int {
	newStack := make([]int, len(deck))
	for i := range deck {
		newStack[(i*increment)%len(deck)] = deck[i]
	}

	return newStack
}
