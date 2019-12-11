package day11

import (
	"fmt"

	"github.com/nlowe/aoc2019/intcode"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const (
	panelColorBlack = 0
	panelColorWhite = 1

	directionNorth = 0
	directionEast  = 1
	directionSouth = 2
	directionWest  = 3
)

type robot struct {
	x int
	y int

	facing int
}

var A = &cobra.Command{
	Use:   "11a",
	Short: "Day 11, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func safeMod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}

func a(challenge *challenge.Input) int {
	_, panelCount := walk(challenge, panelColorBlack)
	return panelCount
}

func walk(program *challenge.Input, startingColor int) (paintedPanels map[int]map[int]int, panelCount int) {
	rob := robot{}
	robotInput := make(chan int, 1)
	controller, output := intcode.NewCPUForProgram(<-program.Lines(), robotInput)

	paintedPanels = map[int]map[int]int{}
	paintedPanels[0] = map[int]int{}
	paintedPanels[0][0] = startingColor

	go controller.Run()

	for {
		color := panelColorBlack
		if row, ok := paintedPanels[rob.x]; ok {
			if c, ok := row[rob.y]; ok {
				color = c
			}
		}

		robotInput <- color

		targetColor, ok := <-output
		if !ok {
			break
		}

		turn, ok := <-output
		if !ok {
			break
		}

		if _, ok := paintedPanels[rob.x]; !ok {
			paintedPanels[rob.x] = map[int]int{}
		}

		if _, alreadyPainted := paintedPanels[rob.x][rob.y]; !alreadyPainted {
			panelCount++
		}

		paintedPanels[rob.x][rob.y] = targetColor

		if turn == 1 {
			rob.facing++
		} else {
			rob.facing--
		}

		rob.facing = safeMod(rob.facing, 4)
		switch rob.facing {
		case directionNorth:
			rob.y++
		case directionEast:
			rob.x++
		case directionSouth:
			rob.y--
		case directionWest:
			rob.x--
		}
	}

	return
}
