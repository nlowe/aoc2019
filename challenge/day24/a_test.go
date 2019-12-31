package day24

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`....#
#..#.
#..##
..#..
#....`)

	result := a(input)

	require.Equal(t, 2129920, result)
}
