package day9

import (
	"fmt"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "9b",
	Short: "Day 9, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) (result int) {
	cpu, outputs := intcode.NewCPUForProgram(<-challenge.Lines(), input.NewFixed(2))

	wg := output.Single(outputs, &result)

	cpu.Run()
	wg.Wait()

	return
}
