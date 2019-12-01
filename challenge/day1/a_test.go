package day1

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`12
14
1969
100756
`)

	result := a(input)

	require.Equal(t, 2+2+654+33583, result)
}
