package day3

import (
	"strconv"
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			expected: 159,
		},
		{
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			expected: 135,
		},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.expected), func(t *testing.T) {
			result := a(challenge.FromLiteral(tt.input))

			require.Equal(t, tt.expected, result)
		})
	}
}
