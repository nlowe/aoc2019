package day20

import (
	"fmt"
	"math"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2019/challenge"
)

const (
	tileOpen       = '.'
	tileWall       = '#'
	tileEmptySpace = ' '

	gateStart = 'A'
	gateEnd   = 'Z'

	// TODO: Re-size the z level dynamically. 128 works for my input and the tests
	//       but may not work for other inputs.
	maxLevels = 128
)

type tile struct {
	x int
	y int
	z int
	r rune

	gateExit *tile
	d        *donut
}

func (t *tile) String() string {
	return fmt.Sprintf("%s {%d,%d,%d}", string(t.r), t.x, t.y, t.z)
}

func (t *tile) isGate() bool {
	return t.r >= 'A' && t.r <= 'Z'
}

func (t *tile) isInnerGate() bool {
	return t.isGate() && t.x > 3 && t.x < (t.d.sx-3) && t.y > 3 && t.y < (t.d.sy-3)
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
		if other := t.d.tileAt(t.x+delta.x, t.y+delta.y, t.z); other != nil {
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
			exit := other.gateExit
			if t.d.gatesChangeLevels {
				if other.z == 0 && !other.isInnerGate() {
					// Outer gates on level 0 are walls
					continue
				}

				if other.isInnerGate() {
					// inner portal, go up
					exit = t.d.tileAt(exit.x, exit.y, exit.z+1)
				} else {
					// outer portal, go down
					exit = t.d.tileAt(exit.x, exit.y, exit.z-1)
				}
			}

			result = append(result, exit)
		} else {
			result = append(result, other)
		}
	}

	return
}

func (t *tile) PathNeighborCost(_ astar.Pather) float64 {
	return 1.0
}

func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	// TODO: Figure out a good cost estimate. If we just use manhattan distance
	//       like in other days, gates that are "far" away may not be taken even
	//       though the path **through** them is shorter. By hard-coding this to
	//       1 we're essentially forcing A* to operate as Dijkstra instead.
	other := to.(*tile)
	return math.Max(0, float64(other.z-1)*64.0)
}

type donut struct {
	d map[int]map[int]map[int]*tile

	sx int
	sy int

	start *tile
	end   *tile

	gatesChangeLevels bool
}

func (d *donut) set(x, y, z int, r rune) {
	if _, ok := d.d[x]; !ok {
		d.d[x] = map[int]map[int]*tile{}
	}

	if _, ok := d.d[x][y]; !ok {
		d.d[x][y] = map[int]*tile{}
	}

	if x > d.sx {
		d.sx = x
	}

	if y > d.sy {
		d.sy = y
	}

	d.d[x][y][z] = &tile{x, y, z, r, nil, d}
}

func (d *donut) tileAt(x, y, z int) *tile {
	if z >= maxLevels {
		panic(fmt.Errorf("tried to access tile at {%d,%d} past max level of %d", x, y, maxLevels))
	}

	if _, ok := d.d[x]; !ok {
		return nil
	}

	if _, ok := d.d[x][y]; !ok {
		return nil
	}

	return d.d[x][y][z]
}

func NewDonut(challenge *challenge.Input) *donut {
	result := &donut{d: map[int]map[int]map[int]*tile{}}
	var startCandidates, endCandidates, gates []*tile

	y := 1
	for line := range challenge.Lines() {
		x := 1
		for _, r := range line {
			for z := 0; z < maxLevels; z++ {
				result.set(x, y, z, r)

				t := result.tileAt(x, y, z)

				if r == 'A' {
					startCandidates = append(startCandidates, t)
				} else if r == 'Z' {
					endCandidates = append(endCandidates, t)
				}

				if t.isGate() {
					gates = append(gates, t)
				}
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
	for _, a := range gates {
		want := a.d.gatePair(a)
		if want == nil {
			continue
		}

		for _, b := range gates {
			if b == a || b == want || b.r != a.r || b.z != a.z {
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
	want := g.d.tileAt(g.x, g.y+1, g.z)
	if want == nil || !want.isGate() || g.r == want.r {
		want = g.d.tileAt(g.x+1, g.y, g.z)
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
