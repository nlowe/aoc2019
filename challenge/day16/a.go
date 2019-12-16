package day16

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2019/util"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "16a",
	Short: "Day 16, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", a(challenge.FromFile()))
	},
}

const iterations = 100

var pattern = [4]int{0, 1, 0, -1}

func a(challenge *challenge.Input) string {
	in := <-challenge.Lines()

	data := make([]int, len(in))
	for i, r := range in {
		data[i] = util.MustAtoI(string(r))
	}

	for i := 0; i < iterations; i++ {
		data = fft(data)
	}

	result := strings.Builder{}
	for i := 0; i < 8; i++ {
		result.WriteString(strconv.Itoa(data[i]))
	}

	return result.String()
}

func fft(in []int) []int {
	result := make([]int, len(in))

	for i := 0; i < len(in); i++ {
		sum := 0
		for j := i; j < len(in); j++ {
			sum += in[j] * mixin(j, i)
		}

		result[i] = int(math.Abs(float64(sum))) % 10
	}

	return result
}

func mixin(offset, iteration int) int {
	return pattern[((1+offset)/(1+iteration))%len(pattern)]
}
