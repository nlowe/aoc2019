package day3

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/cobra"
)

const (
	DirectionUp    = 'U'
	DirectionDown  = 'D'
	DirectionLeft  = 'L'
	DirectionRight = 'R'
)

type segment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (s segment) intersection(other segment) (int, int, bool) {
	sx := float64(s.x2) - float64(s.x1)
	sy := float64(s.y2) - float64(s.y1)
	ox := float64(other.x2) - float64(other.x1)
	oy := float64(other.y2) - float64(other.y1)

	u := (-sy*float64(s.x1-other.x1) + sx*float64(s.y1-other.y1)) / (-ox*sy + sx*oy)
	v := (ox*float64(s.y1-other.y1) - oy*float64(s.x1-other.x1)) / (-ox*sy + sx*oy)

	return s.x1 + int(v*sx), s.y1 + int(v*sy), u >= 0 && u <= 1 && v >= 0 && v <= 1
}

var A = &cobra.Command{
	Use:   "3a",
	Short: "Day 3, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(input *challenge.Input) int {
	wires := input.Lines()
	w1 := makeSegments(<-wires)
	w2 := makeSegments(<-wires)

	var intersections []int

	for _, a := range w1 {
		for _, b := range w2 {
			x, y, intersects := a.intersection(b)
			if d := ManhattanDistance(0, 0, x, y); intersects && d > 0 {
				intersections = append(intersections, d)
			}
		}
	}

	if len(intersections) == 0 {
		panic("no solution")
	}

	sort.Ints(intersections)
	return intersections[0]
}

func makeSegments(steps string) (result []segment) {
	x := 0
	y := 0
	for _, step := range strings.Split(steps, ",") {
		x2 := x
		y2 := y

		magnitude := util.MustAtoI(step[1:])

		switch step[0] {
		case DirectionDown:
			y2 = y - magnitude
		case DirectionUp:
			y2 = y + magnitude
		case DirectionLeft:
			x2 = x - magnitude
		case DirectionRight:
			x2 = x + magnitude
		}

		result = append(result, segment{
			x1: x,
			y1: y,
			x2: x2,
			y2: y2,
		})

		x = x2
		y = y2
	}

	return
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x2-x1))) + int(math.Abs(float64(y2-y1)))
}
