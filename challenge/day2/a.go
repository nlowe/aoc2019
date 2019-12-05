package day2

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "2a",
	Short: "Day 2, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(input *challenge.Input) int {
	cpu, _ := intcode.NewCPUForProgram(<-input.Lines(), nil)

	cpu.Memory[1] = 12
	cpu.Memory[2] = 2

	cpu.Run()
	return cpu.Memory[0]
}
