package day12

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "12b",
	Short: "Day 12, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

func stateEqual(a, b []*moon) (x, y, z bool) {
	x = true
	y = true
	z = true

	for i := range a {
		if a[i].position.x != b[i].position.x || a[i].velocity.x != b[i].velocity.x {
			x = false
		}

		if a[i].position.y != b[i].position.y || a[i].velocity.y != b[i].velocity.y {
			y = false
		}

		if a[i].position.z != b[i].position.z || a[i].velocity.z != b[i].velocity.z {
			z = false
		}
	}

	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func b(challenge *challenge.Input) int {
	var initial []*moon
	var moons []*moon

	for line := range challenge.Lines() {
		initial = append(initial, newMoon(line))
		moons = append(moons, newMoon(line))
	}

	xPhase := 0
	yPhase := 0
	zPhase := 0

	steps := 0
	for xPhase == 0 || yPhase == 0 || zPhase == 0 {
		simStep(moons)
		steps++

		x, y, z := stateEqual(initial, moons)
		if x && xPhase == 0 {
			xPhase = steps
		}

		if y && yPhase == 0 {
			yPhase = steps
		}

		if z && zPhase == 0 {
			zPhase = steps
		}
	}

	return lcm(xPhase, lcm(yPhase, zPhase))
}
