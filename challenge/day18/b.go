package day18

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "18b",
	Short: "Day 18, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) int {
	wholeMaze := ParseMaze(challenge)
	tl, tr, bl, br := split(wholeMaze)

	pruneDeadDoors(tl)
	pruneDeadDoors(tr)
	pruneDeadDoors(bl)
	pruneDeadDoors(br)

	// In general, my specific input can solve this by solving simplified
	// quadrants. This won't work for all inputs so there are no tests.
	// TODO: Find a solution that works for all inputs
	return solve(tl) + solve(tr) + solve(bl) + solve(br)
}

func split(m *maze) (tl, tr, bl, br *maze) {
	tlData := strings.Builder{}
	trData := strings.Builder{}
	blData := strings.Builder{}
	brData := strings.Builder{}

	for y := 0; y <= m.sy; y++ {
		for x := 0; x <= m.sx; x++ {
			t := m.tileAt(x, y)
			if t == nil {
				panic("bad data")
			}

			if x <= m.startX && y <= m.startY {
				if x == m.startX-1 && y == m.startY-1 {
					tlData.WriteRune(tileStart)
				} else if t.x == m.startX || t.y == m.startY {
					tlData.WriteRune(tileWall)
				} else {
					tlData.WriteRune(t.r)
				}
			}

			if x >= m.startX && y <= m.startY {
				if x == m.startX+1 && y == m.startY-1 {
					trData.WriteRune(tileStart)
				} else if t.x == m.startX || t.y == m.startY {
					trData.WriteRune(tileWall)
				} else {
					trData.WriteRune(t.r)
				}
			}

			if x <= m.startX && y >= m.startY {
				if x == m.startX-1 && y == m.startY+1 {
					blData.WriteRune(tileStart)
				} else if t.x == m.startX || t.y == m.startY {
					blData.WriteRune(tileWall)
				} else {
					blData.WriteRune(t.r)
				}
			}

			if x >= m.startX && y >= m.startY {
				if x == m.startX+1 && y == m.startY+1 {
					brData.WriteRune(tileStart)
				} else if t.x == m.startX || t.y == m.startY {
					brData.WriteRune(tileWall)
				} else {
					brData.WriteRune(t.r)
				}
			}
		}

		if y <= m.startY {
			tlData.WriteRune('\n')
			trData.WriteRune('\n')
		}

		if y >= m.startY {
			blData.WriteRune('\n')
			brData.WriteRune('\n')
		}
	}

	tl = ParseMaze(challenge.FromLiteral(tlData.String()))
	tr = ParseMaze(challenge.FromLiteral(trData.String()))
	bl = ParseMaze(challenge.FromLiteral(blData.String()))
	br = ParseMaze(challenge.FromLiteral(brData.String()))

	return
}

func pruneDeadDoors(m *maze) {
	for x, col := range m.m {
		for y, t := range col {
			if t.isDoor() && m.keys['a'+(t.r-'A')] == nil {
				m.m[x][y] = &tile{x, y, tileOpen, m}
			}
		}
	}
}
