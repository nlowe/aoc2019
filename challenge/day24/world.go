package day24

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
)

const (
	width  = 5
	height = 5

	tileEmpty = '.'
	tileBug   = '#'
)

type world struct {
	m [][]bool
}

func emptyWorld() [][]bool {
	m := make([][]bool, width)
	for i := 0; i < width; i++ {
		m[i] = make([]bool, height)
	}

	return m
}

func NewWorld(challenge *challenge.Input) *world {
	result := &world{m: emptyWorld()}

	y := 0
	for row := range challenge.Lines() {
		for x, c := range row {
			if c == tileBug {
				result.m[x][y] = true
			}
		}

		y++
	}

	return result
}

func (w *world) countNeighbors(x, y int) (result int) {
	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		dx := x + delta.x
		dy := y + delta.y

		if dx < 0 || dx >= width || dy < 0 || dy >= height {
			continue
		}

		if w.m[dx][dy] {
			result++
		}
	}

	return
}

func (w *world) tick() {
	next := emptyWorld()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			next[x][y] = w.m[x][y]

			n := w.countNeighbors(x, y)
			if w.m[x][y] && n != 1 {
				next[x][y] = false
			} else if !w.m[x][y] && (n == 1 || n == 2) {
				next[x][y] = true
			}
		}
	}

	w.m = next
}

func (w *world) biodiversity() int {
	pow := 1
	biodiversity := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if w.m[x][y] {
				biodiversity += pow
			}
			pow <<= 1
		}
	}

	return biodiversity
}

func (w *world) dump() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if w.m[x][y] {
				fmt.Print(string(tileBug))
			} else {
				fmt.Print(string(tileEmpty))
			}
		}

		fmt.Println()
	}
}
