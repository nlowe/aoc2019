package day5

import (
	"fmt"
	"sync"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "5a",
	Short: "Day 5, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	cpu, output := intcode.NewCPUForProgram(<-challenge.Lines(), input.NewFixed(1))

	lastCode := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for code := range output {
			fmt.Printf("output: %d\n", code)
			lastCode = code
		}

		wg.Done()
	}()

	cpu.Run()
	wg.Wait()
	return lastCode
}
