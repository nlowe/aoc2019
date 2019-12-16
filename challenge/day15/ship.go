package day15

import (
	"math"

	"github.com/beefsack/go-astar"
)

const (
	statusWall = iota
	statusOk
	statusOxygenSystemFound
	statusUnknown
)

type tile struct {
	x int
	y int
	t int

	s *ship
}

func (t *tile) PathNeighbors() (result []astar.Pather) {
	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if shipTile := t.s.tileAt(t.x+delta.x, t.y+delta.y); shipTile != nil && shipTile.t == statusOk {
			result = append(result, shipTile)
		}
	}

	return
}

func (t *tile) PathNeighborCost(_ astar.Pather) float64 {
	return 1
}

func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	other := to.(*tile)
	return math.Abs(float64(other.x-t.x)) + math.Abs(float64(other.y-t.y))
}

type ship struct {
	m map[int]map[int]*tile

	MinX int
	MinY int
	MaxX int
	MaxY int
}

func (s *ship) set(x, y, t int) {
	if _, ok := s.m[x]; !ok {
		s.m[x] = map[int]*tile{}
	}
	s.m[x][y] = &tile{x, y, t, s}

	if x < s.MinX {
		s.MinX = x
	} else if x > s.MaxX {
		s.MaxX = x
	}

	if y < s.MinY {
		s.MinY = y
	} else if y > s.MaxY {
		s.MaxY = y
	}
}

func (s *ship) tileAt(x, y int) *tile {
	if _, ok := s.m[x]; !ok {
		return nil
	}

	return s.m[x][y]
}

func (s *ship) each(f func(t *tile)) {
	for _, r := range s.m {
		for _, t := range r {
			f(t)
		}
	}
}
