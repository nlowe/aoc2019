package day24

import (
	"fmt"
	"math/bits"

	"github.com/nlowe/aoc2019/challenge"
)

const (
	width  = 5
	height = 5

	tileEmpty = '.'
	tileBug   = '#'
)

const (
	directionNorth int = iota
	directionEast
	directionSouth
	directionWest
)

type world struct {
	active [][]bool
	next   [][]bool

	outerRecursiveCounter func(direction int) int
	innerRecursiveCounter func(direction int) int
}

func emptyWorld() [][]bool {
	m := make([][]bool, width)
	for i := 0; i < width; i++ {
		m[i] = make([]bool, height)
	}

	return m
}

func NewEmptyWorld() *world {
	return &world{active: emptyWorld(), next: emptyWorld()}
}

func NewWorld(challenge *challenge.Input) *world {
	result := NewEmptyWorld()

	y := 0
	for row := range challenge.Lines() {
		for x, c := range row {
			if c == tileBug {
				result.active[x][y] = true
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
			if w.outerRecursiveCounter == nil {
				continue
			}

			if delta.x == -1 {
				result += w.outerRecursiveCounter(directionWest)
			} else if delta.x == 1 {
				result += w.outerRecursiveCounter(directionEast)
			} else if delta.y == -1 {
				result += w.outerRecursiveCounter(directionNorth)
			} else {
				result += w.outerRecursiveCounter(directionSouth)
			}
		} else if dx == width/2 && dy == height/2 && w.innerRecursiveCounter != nil {
			if delta.x == -1 {
				result += w.innerRecursiveCounter(directionEast)
			} else if delta.x == 1 {
				result += w.innerRecursiveCounter(directionWest)
			} else if delta.y == -1 {
				result += w.innerRecursiveCounter(directionSouth)
			} else {
				result += w.innerRecursiveCounter(directionNorth)
			}
		} else if w.active[dx][dy] {
			result++
		}
	}

	return
}

func (w *world) tick() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if w.innerRecursiveCounter != nil &&
				w.outerRecursiveCounter != nil &&
				x == width/2 && y == height/2 {
				continue
			}

			w.next[x][y] = w.active[x][y]

			n := w.countNeighbors(x, y)
			if w.active[x][y] && n != 1 {
				w.next[x][y] = false
			} else if !w.active[x][y] && (n == 1 || n == 2) {
				w.next[x][y] = true
			}
		}
	}
}

func (w *world) swap() {
	w.active = w.next
	w.next = emptyWorld()
}

func (w *world) biodiversity() int {
	pow := 1
	biodiversity := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if w.active[x][y] {
				biodiversity += pow
			}
			pow <<= 1
		}
	}

	return biodiversity
}

func (w *world) count() int {
	return bits.OnesCount32(uint32(w.biodiversity()))
}

func (w *world) dump() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if w.innerRecursiveCounter != nil &&
				w.outerRecursiveCounter != nil &&
				x == width/2 && y == height/2 {
				fmt.Print("?")
			} else if w.active[x][y] {
				fmt.Print(string(tileBug))
			} else {
				fmt.Print(string(tileEmpty))
			}
		}

		fmt.Println()
	}
}
