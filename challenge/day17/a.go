package day17

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "17a",
	Short: "Day 17, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

const (
	tileNewline     = '\n'
	tileEmpty       = '.'
	tileScaffolding = '#'

	robotUp    = '^'
	robotDown  = 'v'
	robotLeft  = '<'
	robotRight = '>'
)

func a(challenge *challenge.Input) int {
	cpu, out := intcode.NewCPUForProgram(<-challenge.Lines(), nil)
	go cpu.Run()

	s := &scaffolding{m: map[int]map[int]*scaffold{}}

	width := 0
	height := 0

	y := 0
	x := 0
	for i := range out {
		if i == tileScaffolding || i == robotLeft || i == robotRight || i == robotUp || i == robotDown {
			s.set(x, y)
		} else if i == tileEmpty {
			s.clear(x, y)
		}

		fmt.Print(string(rune(i)))

		if i == '\n' {
			x = 0
			y++
		} else {
			x++
		}

		if x > width {
			width = x
		}

		if y > height {
			height = y
		}
	}

	alignment := 0
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			if s.isIntersection(x, y) {
				alignment += s.get(x, y).alignment()
			}
		}
	}

	return alignment
}
