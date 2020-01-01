package day24

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(`....#
#..#.
#..##
..#..
#....`)

	iterations = 10
	result := b(input)

	require.Equal(t, 99, result)
}
