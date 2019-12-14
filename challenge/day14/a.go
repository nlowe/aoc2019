package day14

import (
	"fmt"
	"math"
	"strings"

	"github.com/nlowe/aoc2019/util"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "14a",
	Short: "Day 14, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

const (
	componentOre  = "ORE"
	componentFuel = "FUEL"
)

type component struct {
	name  string
	count int
}

type equation struct {
	inputs []component
	output component
}

func parseEquation(eqn string) equation {
	parts := strings.Split(eqn, "=>")
	inputs := strings.Split(parts[0], ",")

	result := equation{}
	for _, input := range inputs {
		parts := strings.Split(strings.TrimSpace(input), " ")
		result.inputs = append(result.inputs, component{
			name:  strings.TrimSpace(parts[1]),
			count: util.MustAtoI(strings.TrimSpace(parts[0])),
		})
	}

	outputParts := strings.Split(strings.TrimSpace(parts[1]), " ")
	result.output = component{
		name:  strings.TrimSpace(outputParts[1]),
		count: util.MustAtoI(strings.TrimSpace(outputParts[0])),
	}

	return result
}

func a(challenge *challenge.Input) int {
	equations := map[string]equation{}
	for line := range challenge.Lines() {
		eqn := parseEquation(line)
		equations[eqn.output.name] = eqn
	}

	return cost(componentFuel, 1, equations, map[string]int{})
}

func cost(name string, n int, equations map[string]equation, leftovers map[string]int) int {
	if name == componentOre {
		return n
	}

	if leftovers[name] >= n {
		leftovers[name] -= n
		return 0
	}

	if leftovers[name] > 0 {
		n -= leftovers[name]
		leftovers[name] = 0
	}

	totalOre := 0
	eqn := equations[name]

	iterationCount := int(math.Ceil(float64(n) / float64(eqn.output.count)))

	for _, c := range eqn.inputs {
		totalOre += cost(c.name, c.count*iterationCount, equations, leftovers)
	}

	leftovers[name] += eqn.output.count*iterationCount - n
	return totalOre
}
