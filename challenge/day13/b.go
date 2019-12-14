package day13

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func init() {
	flags := B.Flags()

	flags.Bool("headless", false, "Run without rendering the game")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func b(challenge *challenge.Input) int {
	// This needs to be buffered because the game doesn't read the paddle state
	// immediately after rendering the ball (it still wants to render more tiles).
	// If we were to block, the render loop wouldn't be able to read more tile data
	// and the CPU would deadlock.
	joystick := make(chan int, 1)
	cpu, out := intcode.NewCPUForProgram(<-challenge.Lines(), joystick)
	cpu.Memory[0] = 2

	headless := viper.GetBool("headless")
	term, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if !headless {
		if err := term.Init(); err != nil {
			panic(err)
		}
		term.DisableMouse()
		defer term.Fini()
	}

	bx := 0
	px := 0

	term.Clear()
	score := 0

	go cpu.Run()
	for {
		x, ok := <-out
		if !ok {
			break
		}

		y, ok := <-out
		if !ok {
			break
		}

		tile, ok := <-out
		if !ok {
			break
		}

		if x == -1 && y == 0 {
			score = tile
			scoreString := fmt.Sprintf("Score: %d", score)
			for i, r := range scoreString {
				term.SetContent(i, 0, r, nil, tcell.StyleDefault)
			}
		} else {
			r := ' '
			switch tile {
			case tileBlock:
				r = '\u2588'
			case tileBall:
				bx = x
				if bx < px {
					joystick <- positionLeft
				} else if bx > px {
					joystick <- positionRight
				} else {
					joystick <- positionNeutral
				}

				r = '\u25ef'
			case tileHorizontalPaddle:
				px = x
				r = '\u2015'
			case tileWall:
				r = '\u2591'
			}

			term.SetContent(x, 1+y, r, nil, tcell.StyleDefault)

			if !headless {
				term.Show()
			}
		}
	}

	return score
}
