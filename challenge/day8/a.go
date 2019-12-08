package day8

import (
	"fmt"
	"math"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

const (
	layerWidth  = 25
	layerHeight = 6
)

var A = &cobra.Command{
	Use:   "8a",
	Short: "Day 8, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	img := <-challenge.Lines()

	counts := [10]int{}

	minZ := math.MaxInt64
	checksum := 0
	layer := 0
	for i, r := range img {
		counts[r-'0']++

		if i%(layerWidth*layerHeight) == 0 && i != 0 {
			if counts[0] < minZ {
				minZ = counts[0]
				checksum = counts[1] * counts[2]
			}
			counts = [10]int{}
			layer++
		}
	}

	return checksum
}
