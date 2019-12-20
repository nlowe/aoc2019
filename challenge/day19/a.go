package day19

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

const (
	width  = 50
	height = 50
)

const (
	statusStationary = iota
	statusAffected
)

var A = &cobra.Command{
	Use:   "19a",
	Short: "Day 19, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	program := <-challenge.Lines()

	affected := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			in := make(chan int)
			cpu, out := intcode.NewCPUForProgram(program, in)
			go cpu.Run()

			in <- x
			in <- y
			status := <-out

			cpu.Halt()

			if status == statusAffected {
				affected++
			}

			switch status {
			case statusStationary:
				fmt.Print(".")
			case statusAffected:
				fmt.Print("#")
			default:
				fmt.Print("?")
			}
		}

		fmt.Println()
	}

	return affected
}
