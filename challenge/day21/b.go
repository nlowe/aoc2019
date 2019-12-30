package day21

import (
	"fmt"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "21b",
	Short: "Day 21, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

const cmdRun = "RUN"

func b(challenge *challenge.Input) int {
	in := make(chan string)
	cpu, out := intcode.NewCPUForProgram(<-challenge.Lines(), input.ASCIIWrapper(in))

	go cpu.Run()

	dmg := 0
	wg := output.Each(out, func(i int) {
		if i < 255 {
			fmt.Print(string(rune(i)))
		}
		dmg = i
	})

	// The springdroid jumps 4 spaces at t time, so there are only 3 cases to check
	// We need to initiate a jump when
	// * the very next tile is a hole (!A)
	// * AND All Of
	//   * There is a hole 3 tiles out (jump as early as possible, !C) AND
	//   * If we jump now we won't fall (D) AND
	//   * If we jump *again* immediately afterwards, we won't fall (H)
	// * OR Both
	//   * The 2nd tile out is a hole that we couldn't make last turn (so !B) AND
	//   * The 4th tile is solid (so we have something to land on, so D)
	// !A || (!C && D && H) || (!B && D)
	in <- "NOT C J"
	in <- "AND D J"
	in <- "AND H J"
	in <- "NOT B T"
	in <- "AND D T"
	in <- "OR T J"
	in <- "NOT A T"
	in <- "OR T J"
	in <- cmdRun

	wg.Wait()
	return dmg
}
