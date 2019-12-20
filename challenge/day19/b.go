package day19

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "19b",
	Short: "Day 19, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

const shipExtraSize = 99

func b(challenge *challenge.Input) int {
	program := <-challenge.Lines()

	lastRowStart := 0
	y := 100
	for {
		x := lastRowStart
		for {
			status := check(x, y, program)

			if status == statusAffected {
				break
			} else {
				x++
			}
		}

		if canFit(x, y, program) {
			return x*10000 + (y - shipExtraSize)
		}

		lastRowStart = x
		y++
	}
}

func check(x, y int, program string) int {
	in := make(chan int)
	cpu, out := intcode.NewCPUForProgram(program, in)
	go cpu.Run()
	defer cpu.Halt()

	in <- x
	in <- y
	return <-out
}

func canFit(x, y int, program string) bool {
	return check(x+shipExtraSize, y-shipExtraSize, program) == statusAffected
}
