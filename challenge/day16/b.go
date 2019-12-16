package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "16b",
	Short: "Day 16, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", b(challenge.FromFile()))
	},
}

func b(challenge *challenge.Input) string {
	in := strings.Repeat(<-challenge.Lines(), 10000)

	data := make([]int, len(in))
	for i, r := range in {
		data[i] = util.MustAtoI(string(r))
	}

	offset := util.MustAtoI(in[:7])
	data = data[offset:]
	for i := 0; i < iterations; i++ {
		data = shortcut(data)
	}

	result := strings.Builder{}
	for i := 0; i < 8; i++ {
		result.WriteString(strconv.Itoa(data[i]))
	}

	return result.String()
}

// TODO: Here be vodo. I need to read up on why this works.
//       this was mostly pieced together from posts on the
//       subreddit.
func shortcut(in []int) []int {
	cumulativeSum := 0
	for _, v := range in {
		cumulativeSum += v
	}

	var result []int
	for i := range in {
		result = append(result, ((cumulativeSum%10)+10)%10)
		cumulativeSum -= in[i]
	}

	return result
}
