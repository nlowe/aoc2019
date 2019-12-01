package day1

import (
	"strconv"
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	testModule := func(mass, expected int) func(*testing.T) {
		return func(t *testing.T) {
			input := challenge.FromLiteral(strconv.Itoa(mass))

			result := b(input)

			require.Equal(t, expected, result)
		}
	}

	t.Run("14", testModule(14, 2))
	t.Run("1969", testModule(1969, 966))
	t.Run("100756", testModule(100756, 50346))
}
