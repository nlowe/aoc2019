package day13

import (
	"fmt"
	"math"

	"github.com/gdamore/tcell"

	"github.com/nlowe/aoc2019/intcode"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const (
	positionNeutral = 0
	positionLeft    = -1
	positionRight   = 1
)

var B = &cobra.Command{
	Use:   "13b",
	Short: "Day 13, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	joystick := positionNeutral
	cpu, out := intcode.NewCPUForProgram(<-challenge.Lines(), nil)
	cpu.Memory[0] = 2

	// TODO: Is there a way to keep input and output channels in sync wihtout this hack?
	cpu.UseFloatingInput(func() int {
		return joystick
	})

	term, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := term.Init(); err != nil {
		panic(err)
	}
	term.DisableMouse()
	defer term.Fini()

	bx := 0
	px := math.MaxInt64

	term.Clear()
	score := 0

	go cpu.Run()
	for {
		x, ok := <-out
		if !ok {
			return score
		}

		y, ok := <-out
		if !ok {
			return score
		}

		tile, ok := <-out
		if !ok {
			return score
		}

		if x == -1 && y == 0 {
			score = tile
		} else {
			switch tile {
			case tileBlock:
				term.SetContent(x, y, tcell.RuneBlock, nil, tcell.StyleDefault)
			case tileBall:
				bx = x

				if px != math.MaxInt64 {
					if bx < px {
						joystick = positionLeft
					} else if bx > px {
						joystick = positionRight
					} else {
						joystick = positionNeutral
					}
				}

				term.SetContent(x, y, '0', nil, tcell.StyleDefault)
			case tileHorizontalPaddle:
				px = x
				term.SetContent(x, y, '\u2582', nil, tcell.StyleDefault)
			case tileWall:
				term.SetContent(x, y, '=', nil, tcell.StyleDefault)
			case tileEmpty:
				term.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			}

			term.Show()
		}
	}

	return score
}
