package intcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"
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
	sut, outputs := NewCPUForProgram("00003,0,00004,0,99", input.NewFixed(42))
	wg := output.Expect(t, outputs, 42)

	sut.Run()
	wg.Wait()
}

func TestCPU_Compare(t *testing.T) {
	tests := []struct {
		program  string
		input    <-chan int
		expected int
	}{
		{program: "3,9,8,9,10,9,4,9,99,-1,8", input: input.NewFixed(8), expected: 1},
		{program: "3,9,8,9,10,9,4,9,99,-1,8", input: input.NewFixed(42), expected: 0},
		{program: "3,9,7,9,10,9,4,9,99,-1,8", input: input.NewFixed(7), expected: 1},
		{program: "3,9,7,9,10,9,4,9,99,-1,8", input: input.NewFixed(42), expected: 0},
		{program: "3,3,1108,-1,8,3,4,3,99", input: input.NewFixed(8), expected: 1},
		{program: "3,3,1108,-1,8,3,4,3,99", input: input.NewFixed(42), expected: 0},
		{program: "3,3,1107,-1,8,3,4,3,99", input: input.NewFixed(7), expected: 1},
		{program: "3,3,1107,-1,8,3,4,3,99", input: input.NewFixed(42), expected: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=%d", tt.program, tt.expected), func(t *testing.T) {
			cpu, outputs := NewCPUForProgram(tt.program, tt.input)
			wg := output.Expect(t, outputs, tt.expected)

			cpu.Run()
			wg.Wait()
		})
	}
}

func TestCPU_Jump(t *testing.T) {
	tests := []struct {
		program  string
		input    <-chan int
		expected int
	}{
		{program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: input.NewFixed(0), expected: 0},
		{program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: input.NewFixed(42), expected: 1},
		{program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: input.NewFixed(0), expected: 0},
		{program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: input.NewFixed(42), expected: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=%d", tt.program, tt.expected), func(t *testing.T) {
			cpu, outputs := NewCPUForProgram(tt.program, tt.input)
			wg := output.Expect(t, outputs, tt.expected)

			cpu.Run()
			wg.Wait()
		})
	}
}

func TestCPU_CmpAndJump(t *testing.T) {
	tests := []struct {
		program  string
		input    <-chan int
		expected int
	}{
		{program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: input.NewFixed(7), expected: 999},
		{program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: input.NewFixed(8), expected: 1000},
		{program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: input.NewFixed(9), expected: 1001},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.expected), func(t *testing.T) {
			cpu, outputs := NewCPUForProgram(tt.program, tt.input)
			wg := output.Expect(t, outputs, tt.expected)

			cpu.Run()
			wg.Wait()
		})
	}
}
