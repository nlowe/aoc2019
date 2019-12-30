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

type strategy struct {
	a int
	b int
}

func (s strategy) compose(other strategy, mod int) strategy {
	return strategy{
		a: safeMod(s.a*other.a, mod),
		b: safeMod(s.b*other.a+other.b, mod),
	}
}

func (s strategy) apply(x, mod int) int {
	return safeMod(s.a*x+s.b, mod)
}

func a(challenge *challenge.Input) int {
	return parseInstructions(challenge, 10007).apply(2019, 10007)
}

func parseInstructions(challenge *challenge.Input, deckSize int) strategy {
	f := strategy{a: 1, b: 0}
	for instruction := range challenge.Lines() {
		if instruction == instrDealNewStack {
			f = f.compose(strategy{a: -1, b: -1}, deckSize)
		} else if strings.HasPrefix(instruction, instrCut) {
			n, err := strconv.Atoi(strings.TrimPrefix(instruction, instrCut))
			if err != nil {
				panic(err)
			}

			f = f.compose(strategy{a: 1, b: -1 * n}, deckSize)
		} else if strings.HasPrefix(instruction, instrDealWithIncrement) {
			increment, err := strconv.Atoi(strings.TrimPrefix(instruction, instrDealWithIncrement))
			if err != nil {
				panic(err)
			}

			f = f.compose(strategy{a: increment, b: 0}, deckSize)
		}
	}

	return f
}

func safeMod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}
