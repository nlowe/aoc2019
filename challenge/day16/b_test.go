package day16

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{in: "03036732577212944063491565474664", out: "84462026"},
		{in: "02935109699940807407585447034323", out: "78725270"},
		{in: "03081770884921959731165446850517", out: "53553731"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			require.Equal(t, tt.out, b(challenge.FromLiteral(tt.in)))
		})
	}
}
