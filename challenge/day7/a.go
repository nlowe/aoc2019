package day7

import (
	"fmt"
	"sync"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "7a",
	Short: "Day 7, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	program := <-challenge.Lines()

	bestValue := 0
	for settings := range phaseGenerator() {
		a, aOut := intcode.NewCPUForProgram(program, input.Prefix(settings[0], input.NewFixed(0)))
		b, bOut := intcode.NewCPUForProgram(program, input.Prefix(settings[1], aOut))
		c, cOut := intcode.NewCPUForProgram(program, input.Prefix(settings[2], bOut))
		d, dOut := intcode.NewCPUForProgram(program, input.Prefix(settings[3], cOut))
		e, eOut := intcode.NewCPUForProgram(program, input.Prefix(settings[4], dOut))

		thrusterValue := 0
		t := output.Single(eOut, &thrusterValue)

		wg := sync.WaitGroup{}
		wg.Add(5)
		for _, amp := range []*intcode.CPU{a, b, c, d, e} {
			go func(c *intcode.CPU) {
				defer wg.Done()
				c.Run()
			}(amp)
		}

		t.Wait()
		wg.Wait()

		if thrusterValue > bestValue {
			fmt.Printf("New Best thruster value: %d, Phases: %+v\n", thrusterValue, settings)
			bestValue = thrusterValue
		}
	}

	return bestValue
}

// Generate all possible phases using Heap's algorithm
func phaseGenerator() <-chan []int {
	result := make(chan []int)

	go func() {
		permute(5, []int{0, 1, 2, 3, 4}, result)
		close(result)
	}()
	return result
}

func permute(n int, input []int, output chan<- []int) {
	if n == 1 {
		v := make([]int, len(input))
		copy(v, input)
		output <- v
		return
	}

	for i := 0; i < n; i++ {
		permute(n-1, input, output)

		if n%2 == 0 {
			input[i], input[n-1] = input[n-1], input[i]
		} else {
			input[0], input[n-1] = input[n-1], input[0]
		}
	}
}
