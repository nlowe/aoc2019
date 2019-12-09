package day9

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := b(input)

	require.Equal(t, 42, result)
}
