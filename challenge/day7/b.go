package day7

import (
	"fmt"
	"sync"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "7b",
	Short: "Day 7, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	program := <-challenge.Lines()

	bestValue := 0
	for settings := range phaseGenerator([]int{5, 6, 7, 8, 9}) {
		feedback := make(chan int)

		a, aOut := intcode.NewCPUForProgram(program, input.Prefix(settings[0], input.Prefix(0, feedback)))
		b, bOut := intcode.NewCPUForProgram(program, input.Prefix(settings[1], aOut))
		c, cOut := intcode.NewCPUForProgram(program, input.Prefix(settings[2], bOut))
		d, dOut := intcode.NewCPUForProgram(program, input.Prefix(settings[3], cOut))
		e, eOut := intcode.NewCPUForProgram(program, input.Prefix(settings[4], dOut))

		thrusterValue := 0
		t := output.Each(eOut, func(v int) {
			thrusterValue = v
			feedback <- v
		})

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
		close(feedback)

		if thrusterValue > bestValue {
			fmt.Printf("New Best thruster value: %d, Phases: %+v\n", thrusterValue, settings)
			bestValue = thrusterValue
		}
	}

	return bestValue
}
