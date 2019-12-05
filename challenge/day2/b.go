package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

const target = 19690720

var B = &cobra.Command{
	Use:   "2b",
	Short: "Day 2, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(input *challenge.Input) int {
	program := <-input.Lines()

	for verb := 0; verb <= 99; verb++ {
		for noun := 0; noun <= 99; noun++ {
			cpu, _ := intcode.NewCPUForProgram(patch(program, verb, noun), nil)

			cpu.Run()

			if cpu.Memory[0] == target {
				return 100*noun + verb
			}
		}
	}

	panic("no solution found")
}

func patch(program string, verb, noun int) string {
	parts := strings.Split(program, ",")
	parts[1] = strconv.Itoa(noun)
	parts[2] = strconv.Itoa(verb)

	return strings.Join(parts, ",")
}
