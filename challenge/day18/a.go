package day18

import (
	"fmt"
	"math"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "18a",
	Short: "Day 18, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

type cacheKey struct {
	x    int
	y    int
	have keyset
}

type searchState struct {
	x    int
	y    int
	have keyset
	cost int
}

func a(challenge *challenge.Input) int {
	m := ParseMaze(challenge)

	return solve(m)
}

func solve(m *maze) int {
	all := keyset(0)
	for k := range m.keys {
		all.add(k)
	}

	start := m.tileAt(m.startX, m.startY)
	cache := map[cacheKey]int{}

	workQueue := []searchState{{start.x, start.y, NoKeys, 0}}
	for len(workQueue) > 0 {
		var work searchState
		work, workQueue = workQueue[0], workQueue[1:]

		tile := m.tileAt(work.x, work.y)
		if tile.isKey() {
			work.have.add(tile.r)
			if work.have == all {
				return work.cost
			}
		} else if tile.isDoor() && !work.have.has('a'+(tile.r-'A')) {
			continue
		}

		cc := cacheKey{work.x, work.y, work.have}
		costSoFar, found := cache[cc]
		if !found {
			costSoFar = math.MaxInt64
		}

		if work.cost >= costSoFar {
			continue
		}

		cache[cc] = work.cost
		work.cost++

		for _, delta := range []struct {
			x int
			y int
		}{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			if t := m.tileAt(work.x+delta.x, work.y+delta.y); t != nil && t.r != tileWall {
				workQueue = append(workQueue, searchState{t.x, t.y, work.have, work.cost})
			}
		}
	}

	panic("no solution")
}
