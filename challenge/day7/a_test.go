package day7

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestPermute(t *testing.T) {
	sut := []int{0, 1, 2}
	permutations := make(chan []int)

	var generated [][]int
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for p := range permutations {
			fmt.Printf("got %+v\n", p)
			generated = append(generated, p)
		}

		wg.Done()
	}()

	permute(len(sut), sut, permutations)
	close(permutations)
	wg.Wait()

	require.Len(t, generated, 3*2*1)
	assert.Contains(t, generated, []int{0, 1, 2})
	assert.Contains(t, generated, []int{0, 2, 1})
	assert.Contains(t, generated, []int{1, 0, 2})
	assert.Contains(t, generated, []int{1, 2, 0})
	assert.Contains(t, generated, []int{2, 0, 1})
	assert.Contains(t, generated, []int{2, 1, 0})
}

func TestA(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", expected: 43210},
		{input: "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", expected: 54321},
		{input: "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", expected: 65210},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := challenge.FromLiteral(tt.input)

			result := a(input)
			require.Equal(t, tt.expected, result)
		})
	}
}
