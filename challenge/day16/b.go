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
	if offset < len(data)/2 {
		panic(fmt.Errorf("shortcut won't work for input of length %d (wanted offset %d)", len(data), offset))
	}

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

// shortcut performs a fast FFT with the assumption that the input is second half of the actual input
//
// Let's assume our input length is 20 and the answer offset is at least at position 10 or later. The matrix formed by
// the phase factors has the following format (split into quadrants for visibility):
//
//     1  0 -1  0  1  0 -1  0  1  0   |  -1  0  1  0 -1  0  1  0 -1  0
//     0  1  1  0  0 -1 -1  0  0  1   |   1  0  0 -1 -1  0  0  1  1  0
//     0  0  1  1  1  0  0  0 -1 -1   |  -1  0  0  0  1  1  1  0  0  0
//     0  0  0  1  1  1  1  0  0  0   |   0 -1 -1 -1 -1  0  0  0  0  1
//     0  0  0  0  1  1  1  1  1  0   |   0  0  0  0 -1 -1 -1 -1 -1  0
//     0  0  0  0  0  1  1  1  1  1   |   1  0  0  0  0  0  0 -1 -1 -1
//     0  0  0  0  0  0  1  1  1  1   |   1  1  1  0  0  0  0  0  0  0
//     0  0  0  0  0  0  0  1  1  1   |   1  1  1  1  1  0  0  0  0  0
//     0  0  0  0  0  0  0  0  1  1   |   1  1  1  1  1  1  1  0  0  0
//     0  0  0  0  0  0  0  0  0  1   |   1  1  1  1  1  1  1  1  1  0
//                                    |
//     -------------------------------+--------------------------------
//                                    |
//     0  0  0  0  0  0  0  0  0  0   |   1  1  1  1  1  1  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  1  1  1  1  1  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  1  1  1  1  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  1  1  1  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  0  1  1  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  0  0  1  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  0  0  0  1  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  0  0  0  0  1  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  0  0  0  0  0  1  1
//     0  0  0  0  0  0  0  0  0  0   |   0  0  0  0  0  0  0  0  0  1
//
// Remember that row i in the matrix modifies digit i in each iteration. Since our offset is at least within the second
// half we only care about the **bottom** half of the phase transform matrix, as these are the only rows that contribute
// to the result.
//
// Additionally, notice that the only phase factor that we're multiplying by is either 1 or 0 by this point. The FFT
// calculates the value of each digit by multiplying each of the input digits pair-wise by each element from a row of a
// column and then taking the sum of all of those calculations mod10. Since the left side of the bottom half of the
// transform matrix is all zero, none of those digits contribute to the result in any iteration either, so we can
// discard them.
//
// Let's zoom in on the quadrant that we care about:
//
//     1  1  1  1  1  1  1  1  1  1
//     0  1  1  1  1  1  1  1  1  1
//     0  0  1  1  1  1  1  1  1  1
//     0  0  0  1  1  1  1  1  1  1
//     0  0  0  0  1  1  1  1  1  1
//     0  0  0  0  0  1  1  1  1  1
//     0  0  0  0  0  0  1  1  1  1
//     0  0  0  0  0  0  0  1  1  1
//     0  0  0  0  0  0  0  0  1  1
//     0  0  0  0  0  0  0  0  0  1
//
// Recall from earlier that when we see a 0 in this phase transform matrix we can discard that digit from the
// calculation each iteration. Simply put, this means that digit 0 is the sum of all digits mod10, digit 1 is the sum
// of digits 1..n mod10, digit 2 is the sum of digits 2..n mod10, and so on. This means we can quickly calculate the
// value of all digits by first taking the cumulative sum of all digits, calculating the cumulative sum mod10, and
// finally subtracting the initial digit from the cumulative sum (for the next iteration since this digit will no longer
// contribute to the result for the rest of the iteration).
func shortcut(in []int) []int {
	cumulativeSum := 0
	for _, v := range in {
		cumulativeSum += v
	}

	var result []int
	for i := range in {
		result = append(result, cumulativeSum%10)
		cumulativeSum -= in[i]
	}

	return result
}
