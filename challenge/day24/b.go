package day24

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "24b",
	Short: "Day 24, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

var iterations = 200

type recursiveWorld struct {
	levels map[int]*world
}

func makeRecursive(w *world) *recursiveWorld {
	result := &recursiveWorld{levels: map[int]*world{0: w}}

	for i := -1; i >= (-1*iterations/2)-1; i-- {
		result.createLevel(i)
	}

	for i := 1; i <= iterations/2+1; i++ {
		result.createLevel(i)
	}

	w.innerRecursiveCounter = result.makeInnerRecursiveCounter(0)
	w.outerRecursiveCounter = result.makeOuterRecursiveCounter(0)
	return result
}

func (r *recursiveWorld) createLevel(level int) {
	w := NewEmptyWorld()
	w.innerRecursiveCounter = r.makeInnerRecursiveCounter(level)
	w.outerRecursiveCounter = r.makeOuterRecursiveCounter(level)
	r.levels[level] = w
}

func (r *recursiveWorld) tickAll() {
	// tick all levels
	for _, w := range r.levels {
		w.tick()
	}
}

func (r *recursiveWorld) swapAll() {
	for _, w := range r.levels {
		w.swap()
	}
}

func (r *recursiveWorld) popCount() (result int) {
	for _, w := range r.levels {
		result += w.count()
	}

	return
}

func (r *recursiveWorld) makeOuterRecursiveCounter(level int) func(int) int {
	return func(direction int) int {
		outer, ok := r.levels[level-1]
		if !ok {
			return 0
		}

		switch direction {
		case directionNorth:
			if outer.active[2][1] {
				return 1
			}
		case directionEast:
			if outer.active[3][2] {
				return 1
			}
		case directionSouth:
			if outer.active[2][3] {
				return 1
			}
		default:
			if outer.active[1][2] {
				return 1
			}
		}

		return 0
	}
}

func (r *recursiveWorld) makeInnerRecursiveCounter(level int) func(int) int {
	return func(direction int) (result int) {
		inner, ok := r.levels[level+1]
		if !ok {
			return 0
		}

		switch direction {
		case directionNorth:
			for x := 0; x < width; x++ {
				if inner.active[x][0] {
					result++
				}
			}
		case directionSouth:
			for x := 0; x < width; x++ {
				if inner.active[x][height-1] {
					result++
				}
			}
		case directionEast:
			for y := 0; y < height; y++ {
				if inner.active[width-1][y] {
					result++
				}
			}
		case directionWest:
			for y := 0; y < height; y++ {
				if inner.active[0][y] {
					result++
				}
			}
		}

		return
	}
}

func b(challenge *challenge.Input) int {
	r := makeRecursive(NewWorld(challenge))

	for i := 0; i < iterations; i++ {
		r.tickAll()
		r.swapAll()
	}

	return r.popCount()
}
