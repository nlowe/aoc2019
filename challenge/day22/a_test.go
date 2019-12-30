package day22

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestFancyShuffle(t *testing.T) {
	deckSize = 10
	for _, tt := range []struct {
		instructions string
		expected     []int
	}{
		{instructions: "deal into new stack", expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{instructions: "cut 3", expected: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}},
		{instructions: "cut -4", expected: []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}},
		{instructions: "deal with increment 3", expected: []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}},
		{instructions: `deal with increment 7
deal into new stack
deal into new stack`, expected: []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}},
		{instructions: `cut 6
deal with increment 7
deal into new stack`, expected: []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}},
		{instructions: `deal with increment 7
deal with increment 9
cut -2`, expected: []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}},
		{instructions: `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`, expected: []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}},
	} {
		t.Run(tt.instructions, func(t *testing.T) {
			require.Equal(t, tt.expected, fancyShuffle(challenge.FromLiteral(tt.instructions)))
		})
	}
}
