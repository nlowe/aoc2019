package day8

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "8b",
	Short: "Day 8, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Answer:")
		b(challenge.FromFile())
		fmt.Println()
	},
}

func b(challenge *challenge.Input) {
	img := <-challenge.Lines()

	data := [layerWidth * layerHeight]rune{}
	for i := range data {
		data[i] = '2'
	}

	for i, r := range img {
		idx := i % (layerWidth * layerHeight)
		if r != '2' && data[idx] == '2' {
			data[idx] = r
		}
	}

	for i, r := range data {
		if i%layerWidth == 0 && i != 0 {
			fmt.Println()
		}

		switch r {
		case '0':
			fallthrough
		case '2':
			fmt.Print(" ")
		case '1':
			fmt.Print("\u2588")
		}
	}
}
