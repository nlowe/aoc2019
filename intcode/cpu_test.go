package intcode

import (
	"sync"
	"testing"

	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/stretchr/testify/require"
)

func TestCPU_BasicOpCodes(t *testing.T) {
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
			sut, _ := NewCPUForProgram(tt.input, nil)
			sut.Run()

			require.Equal(t, tt.expected, sut.Memory)
		})
	}
}

func TestCPU_Indirection(t *testing.T) {
	sut, _ := NewCPUForProgram("01101,4,1,0,99", nil)
	sut.Run()

	require.Equal(t, []int{5, 4, 1, 0, 99}, sut.Memory)
}

func TestCPU_IO(t *testing.T) {
	sut, output := NewCPUForProgram("00003,0,00004,0,99", input.NewFixed(42))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		require.Equal(t, 42, <-output)
		wg.Done()
	}()

	sut.Run()
	wg.Wait()
}
