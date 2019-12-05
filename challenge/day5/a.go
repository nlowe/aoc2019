package day5

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"
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
	cpu, outputs := intcode.NewCPUForProgram(<-challenge.Lines(), input.NewFixed(1))

	lastCode := 0
	wg := output.Each(outputs, func(code int) {
		fmt.Printf("output: %d\n", code)
		lastCode = code
	})

	cpu.Run()
	wg.Wait()
	return lastCode
}
