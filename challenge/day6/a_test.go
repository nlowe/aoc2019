package day6

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`)

	result := a(input)

	require.Equal(t, 42, result)
}
