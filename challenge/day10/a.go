package day10

import (
	"fmt"
	"math"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const symAsteroid = '#'

var A = &cobra.Command{
	Use:   "10a",
	Short: "Day 10, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

type asteroid struct {
	x float64
	y float64
}

func (a asteroid) distanceTo(other asteroid) float64 {
	return math.Sqrt(math.Pow(a.x-other.x, 2) + math.Pow(a.y-other.y, 2))
}

func (a asteroid) angleTo(other asteroid) float64 {
	theta := math.Atan2(other.x-a.x, a.y-other.y)

	if theta < 0 {
		return theta + 2*math.Pi
	}

	return theta
}

func (a asteroid) slopeTo(other asteroid) float64 {
	if a.y == other.y {
		return math.Inf(1)
	} else if a.x == other.x {
		return 0
	}

	return math.Abs(a.x-other.x) / math.Abs(a.y-other.y)
}

func (a asteroid) canSee(other asteroid, all []asteroid) bool {
	if a.x == other.x && a.y == other.y {
		return false
	}

	for _, blocker := range all {
		if (blocker.x == a.x || blocker.x == other.x) && (blocker.y == a.y || blocker.y == other.y) {
			continue
		} else if blocker.x < math.Min(a.x, other.x) || blocker.x > math.Max(a.x, other.x) {
			continue
		} else if blocker.y < math.Min(a.y, other.y) || blocker.y > math.Max(a.y, other.y) {
			continue
		}

		if a.slopeTo(blocker) == a.slopeTo(other) {
			return a.distanceTo(other) < a.distanceTo(blocker)
		}
	}

	return true
}

func a(challenge *challenge.Input) int {
	asteroids := makeMap(challenge)

	_, result, _ := findStation(asteroids)
	return result
}

func findStation(asteroids []asteroid) (station asteroid, best int, targets []asteroid) {
	for _, a := range asteroids {
		seen := 0
		var inSight []asteroid
		for _, other := range asteroids {
			if a.canSee(other, asteroids) {
				inSight = append(inSight, other)
				seen++
			}
		}

		if seen > best {
			station = a
			best = seen
			targets = inSight
		}
	}

	return
}

func makeMap(challenge *challenge.Input) []asteroid {
	var asteroids []asteroid

	y := 0
	for line := range challenge.Lines() {
		for x, v := range line {
			if v == symAsteroid {
				asteroids = append(asteroids, asteroid{float64(x), float64(y)})
			}
		}

		y++
	}

	return asteroids
}
