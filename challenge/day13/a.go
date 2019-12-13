package day13

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

const (
	tileEmpty = iota
	tileWall
	tileBlock
	tileHorizontalPaddle
	tileBall
)

var A = &cobra.Command{
	Use:   "13a",
	Short: "Day 13, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	cpu, out := intcode.NewCPUForProgram(<-challenge.Lines(), nil)

	screen := map[int]map[int]int{}

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

		if _, ok := screen[x]; !ok {
			screen[x] = map[int]int{}
		}

		screen[x][y] = tile
	}

	tiles := 0
	for _, col := range screen {
		for _, tile := range col {
			if tile == tileBlock {
				tiles++
			}
		}
	}

	return tiles
}
