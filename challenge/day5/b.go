package day5

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "5b",
	Short: "Day 5, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	cpu, outputs := intcode.NewCPUForProgram(<-challenge.Lines(), input.NewFixed(5))

	code := 0
	wg := output.Single(outputs, &code)

	cpu.Run()
	wg.Wait()
	return code
}
