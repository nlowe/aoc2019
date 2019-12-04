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

func TestRepeatRule(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{input: 112233, expected: true},
		{input: 123444, expected: false},
		{input: 111122, expected: true},
		{input: 122333, expected: true},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			require.Equal(t, tt.expected, repeatRule(tt.input))
		})
	}
}
