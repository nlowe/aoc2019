package day11

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "11b",
	Short: "Day 11, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: \n%s\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) string {
	paintedPanels, _ := walk(challenge, panelColorWhite)

	// I don't know if they're all flipped, but these bounds work for me
	result := strings.Builder{}
	for y := 0; y >= -6; y-- {
		for x := 0; x <= 40; x++ {
			color := panelColorBlack
			if row, ok := paintedPanels[x]; ok {
				if c, ok := row[y]; ok {
					color = c
				}
			}

			if color == panelColorWhite {
				result.WriteString("\u2588\u2588")
			} else {
				result.WriteString("  ")
			}
		}

		result.WriteString("\n")
	}

	return result.String()
}
