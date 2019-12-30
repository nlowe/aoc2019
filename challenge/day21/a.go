package day21

import (
	"fmt"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "21a",
	Short: "Day 21, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

const cmdWalk = "WALK"

func a(challenge *challenge.Input) int {
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

	// The springdroid jumps 4 spaces at t time
	// We need to initiate a jump when
	// * the very next tile is a hole (!A) OR
	// * Both
	//   * The 3rd tile out is a hole (jump as early as possible, so !C) AND
	//   * The 4th tile is solid (so we have something to land on, so D)
	// !A || (!C && D)
	in <- "NOT C J"
	in <- "AND D J"
	in <- "NOT A T"
	in <- "OR T J"
	in <- cmdWalk

	wg.Wait()
	return dmg
}
