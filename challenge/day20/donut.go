package day20

import (
	"fmt"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2019/challenge"
)

const (
	tileOpen       = '.'
	tileWall       = '#'
	tileEmptySpace = ' '

	gateStart = 'A'
	gateEnd   = 'Z'
)

type tile struct {
	x int
	y int
	r rune

	gateExit *tile
	d        *donut
}

func (t *tile) String() string {
	return fmt.Sprintf("%s {%d,%d}", string(t.r), t.x, t.y)
}

func (t *tile) isGate() bool {
	return t.r >= 'A' && t.r <= 'Z'
}

func (t *tile) directNeighbors() (result []*tile) {
	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if other := t.d.tileAt(t.x+delta.x, t.y+delta.y); other != nil {
			result = append(result, other)
		}
	}

	return
}

func (t *tile) PathNeighbors() (result []astar.Pather) {
	for _, other := range t.directNeighbors() {
		if other.r == tileWall || other.r == tileEmptySpace {
			continue
		}

		if t == t.d.start && other.r == gateStart {
			continue
		} else if t == t.d.end && other.r == gateEnd {
			continue
		}

		if other.isGate() && other.gateExit != nil {
			result = append(result, other.gateExit)
		} else {
			result = append(result, other)
		}
	}

	return
}

func (t *tile) PathNeighborCost(_ astar.Pather) float64 {
	return 1.0
}

func (t *tile) PathEstimatedCost(_ astar.Pather) float64 {
	// TODO: Figure out a good cost estimate. If we just use manhattan distance
	//       like in other days, gates that are "far" away may not be taken even
	//       though the path **through** them is shorter. By hard-coding this to
	//       1 we're essentially forcing A* to operate as Dijkstra instead.
	return 1.0
}

type donut struct {
	d map[int]map[int]*tile

	start *tile
	end   *tile
}

func (d *donut) set(x, y int, r rune) {
	if _, ok := d.d[x]; !ok {
		d.d[x] = map[int]*tile{}
	}
	d.d[x][y] = &tile{x, y, r, nil, d}
}

func (d *donut) tileAt(x, y int) *tile {
	if _, ok := d.d[x]; !ok {
		return nil
	}

	return d.d[x][y]
}

func NewDonut(challenge *challenge.Input) *donut {
	result := &donut{d: map[int]map[int]*tile{}}
	var startCandidates, endCandidates, gates []*tile

	y := 1
	for line := range challenge.Lines() {
		x := 1
		for _, r := range line {
			result.set(x, y, r)
			t := result.tileAt(x, y)

			if r == 'A' {
				startCandidates = append(startCandidates, t)
			} else if r == 'Z' {
				endCandidates = append(endCandidates, t)
			}

			if t.isGate() {
				gates = append(gates, t)
			}

			x++
		}

		y++
	}

	result.start = findStartOrEnd(startCandidates, gateStart)
	result.end = findStartOrEnd(endCandidates, gateEnd)
	wireGates(gates)

	return result
}

func findStartOrEnd(entranceCandidates []*tile, search rune) *tile {
	for _, e := range entranceCandidates {
		hasMatchingEntrance := false
		var candidate *tile

		for _, ex := range e.directNeighbors() {
			if ex.r == search {
				hasMatchingEntrance = true
			}

			if ex.r == tileOpen {
				candidate = ex
			}
		}

		if hasMatchingEntrance && candidate != nil {
			return candidate
		}
	}

	return nil
}

func wireGates(gates []*tile) {
	fmt.Printf("Have %d partial gates to wire up\n", len(gates))
	for _, a := range gates {
		want := a.d.gatePair(a)
		if want == nil {
			continue
		}

		for _, b := range gates {
			if b == a || b == want || b.r != a.r {
				continue
			}

			got := b.d.gatePair(b)
			if got == nil || got.r != want.r {
				continue
			}

			aSide := gateOpening(a, want)
			bSide := gateOpening(b, got)

			a.gateExit = bSide
			want.gateExit = bSide
			b.gateExit = aSide
			got.gateExit = aSide
		}
	}
}

func (d *donut) gatePair(g *tile) *tile {
	want := g.d.tileAt(g.x, g.y+1)
	if want == nil || !want.isGate() || g.r == want.r {
		want = g.d.tileAt(g.x+1, g.y)
		if want == nil || !want.isGate() || g.r == want.r {
			return nil
		}
	}

	return want
}

func gateOpening(a, b *tile) *tile {
	for _, t := range a.directNeighbors() {
		if t.r == tileOpen {
			return t
		}
	}

	for _, t := range b.directNeighbors() {
		if t.r == tileOpen {
			return t
		}
	}

	panic("could not find gate entrance/exit")
}
