package day12

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`)

	result := b(input)

	require.Equal(t, 2772, result)
}

func TestB_ReallyLong(t *testing.T) {
	input := challenge.FromLiteral(`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`)

	result := b(input)

	require.Equal(t, 4686774924, result)
}
