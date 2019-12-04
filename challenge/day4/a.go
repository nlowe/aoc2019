package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/cobra"
)

const maxPasswordLength = 6

var A = &cobra.Command{
	Use:   "4a",
	Short: "Day 4, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(input *challenge.Input) (count int) {
	low, high := parseRange(input)

	for pw := low; pw <= high; pw++ {
		if isValidPassword(pw) {
			count++
		}
	}

	return
}

func isValidPassword(pw int) bool {
	s := strconv.Itoa(pw)

	// Rule 1: Must be six digits
	if len(s) != maxPasswordLength {
		return false
	}

	// Rule 2: Must be in range (implicit)
	// Rule 3: At least two adjacent digits are the same
	rule3 := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			rule3 = true
			break
		}
	}

	if !rule3 {
		return false
	}

	// Rule 4: The numbers must strictly increase
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] < s[i] {
			return false
		}
	}

	return true
}

func parseRange(input *challenge.Input) (int, int) {
	r := <-input.Lines()
	parts := strings.Split(r, "-")

	return util.MustAtoI(parts[0]), util.MustAtoI(parts[1])
}
