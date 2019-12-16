package day16

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{in: "80871224585914546619083218645595", out: "24176176"},
		{in: "19617804207202209144916044189917", out: "73745418"},
		{in: "69317163492948606335995924319873", out: "52432133"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			require.Equal(t, tt.out, a(challenge.FromLiteral(tt.in)))
		})
	}
}
