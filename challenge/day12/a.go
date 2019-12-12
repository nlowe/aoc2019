package day12

import (
	"fmt"
	"math"
	"strings"

	"github.com/nlowe/aoc2019/util"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "12a",
	Short: "Day 12, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

var steps = 1000

type vec3 struct {
	x int
	y int
	z int
}

func (v vec3) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", v.x, v.y, v.z)
}

type moon struct {
	position vec3
	velocity vec3
}

func newMoon(line string) *moon {
	m := strings.Trim(line, "<>")
	parts := strings.Split(m, ",")

	x := util.MustAtoI(strings.Split(parts[0], "=")[1])
	y := util.MustAtoI(strings.Split(parts[1], "=")[1])
	z := util.MustAtoI(strings.Split(parts[2], "=")[1])

	return &moon{
		position: vec3{
			x: x,
			y: y,
			z: z,
		},
		velocity: vec3{},
	}
}

func intAbs(i int) int {
	return int(math.Abs(float64(i)))
}

func (m *moon) potentialEnergy() int {
	return intAbs(m.position.x) + intAbs(m.position.y) + intAbs(m.position.z)
}

func (m *moon) kineticEnergy() int {
	return intAbs(m.velocity.x) + intAbs(m.velocity.y) + intAbs(m.velocity.z)
}

func (m *moon) String() string {
	return fmt.Sprintf("pos=%s, vel=%s; pot=%d; kin=%d", m.position, m.velocity, m.potentialEnergy(), m.kineticEnergy())
}

func (m *moon) applyGravity(other *moon) {
	if m.position.x > other.position.x {
		m.velocity.x--
		other.velocity.x++
	} else if m.position.x < other.position.x {
		m.velocity.x++
		other.velocity.x--
	}

	if m.position.y > other.position.y {
		m.velocity.y--
		other.velocity.y++
	} else if m.position.y < other.position.y {
		m.velocity.y++
		other.velocity.y--
	}

	if m.position.z > other.position.z {
		m.velocity.z--
		other.velocity.z++
	} else if m.position.z < other.position.z {
		m.velocity.z++
		other.velocity.z--
	}
}

func simStep(moons []*moon) {
	for i := range moons {
		for j := i + 1; j < len(moons); j++ {
			moons[i].applyGravity(moons[j])
		}
	}

	for _, m := range moons {
		m.position.x += m.velocity.x
		m.position.y += m.velocity.y
		m.position.z += m.velocity.z
	}
}

func a(challenge *challenge.Input) int {
	var moons []*moon

	for line := range challenge.Lines() {
		moons = append(moons, newMoon(line))
	}

	for i := 0; i < steps; i++ {
		simStep(moons)
	}

	totalEnergy := 0
	for _, m := range moons {
		totalEnergy += m.kineticEnergy() * m.potentialEnergy()
	}

	return totalEnergy
}
