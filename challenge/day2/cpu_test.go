package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCPU(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "1,9,10,3,2,3,11,0,99,30,40,50",
			expected: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			input:    "1,0,0,0,99",
			expected: []int{2, 0, 0, 0, 99},
		},
		{
			input:    "2,3,0,3,99",
			expected: []int{2, 3, 0, 6, 99},
		},
		{
			input:    "2,4,4,5,99,0",
			expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			input:    "1,1,1,4,99,5,6,0,99",
			expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			sut := NewCPUForProgram(tt.input)
			sut.Run()

			require.Equal(t, tt.expected, sut.Memory)
		})
	}
}