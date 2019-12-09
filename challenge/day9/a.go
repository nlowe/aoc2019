package day9

import (
	"fmt"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "9a",
	Short: "Day 9, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) (result int) {
	cpu, outputs := intcode.NewCPUForProgram(<-challenge.Lines(), input.NewFixed(1))

	wg := output.Single(outputs, &result)

	cpu.Run()
	wg.Wait()

	return
}
