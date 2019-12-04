package day4

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		input         int
		shouldBeValid bool
	}{
		{input: 111111, shouldBeValid: true},
		{input: 223450, shouldBeValid: false},
		{input: 123789, shouldBeValid: false},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			require.Equal(t, tt.shouldBeValid, isValidPassword(tt.input))
		})
	}
}
