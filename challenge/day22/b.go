package day22

import (
	"fmt"
	"math/big"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "22b",
	Short: "Day 22, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

const (
	bigDeckSize int64 = 119315717514047
	iterations  int64 = 101741582076661
)

func (s strategy) pow(n, mod int) strategy {
	m := big.NewInt(int64(mod))
	r := big.NewInt(0).Exp(big.NewInt(int64(1-s.a)), big.NewInt(int64(n-2)), m)
	r.Mul(r, big.NewInt(int64(s.b)))
	r.Mod(r, m)

	return strategy{
		a: 0,
		b: int(r.Int64()),
	}
}

func (s strategy) inv(mod int) strategy {
	return strategy{
		a: safeMod(1/s.a, mod),
		b: safeMod(-s.b/s.a, mod),
	}
}

func b(challenge *challenge.Input) int {
	f := parseInstructions(challenge, int(bigDeckSize))
	a := big.NewInt(int64(f.a))
	b := big.NewInt(int64(f.b))
	m := big.NewInt(bigDeckSize)
	k := big.NewInt(iterations)

	// Most of this is based off of https://codeforces.com/blog/entry/72593
	// Essentially, we compose f with itself k times. This forms a geometric
	// series in the form:
	//
	//         b * (1 - a^k)
	// a^k*x + ------------- mod m
	//             1 - a
	//
	// We can do division by multiplying by the multiplicative inverse
	// of 1/(1-a) mod m.

	fA := &big.Int{}
	fA.Exp(a, k, m)

	fB := &big.Int{}
	top := &big.Int{}
	bottom := &big.Int{}
	bottom.Sub(big.NewInt(1), a)
	fB.Mul(b, top.Sub(big.NewInt(1), fA))
	fB.Mul(fB, bottom.ModInverse(bottom, m))

	// Since part 2 asks for what card is in a given position, we need to
	// take the inverse of this function, which works out to
	//
	// x - fB
	// ------ mod M
	//   fA
	//
	// Where x is the card we are interested in. Division is again performed
	// by multiplying the multiplicative inverse.

	answer := &big.Int{}
	answer.Sub(big.NewInt(2020), fB)
	answer.Mul(answer, fA.ModInverse(fA, m))
	answer.Mod(answer, m)

	return int(answer.Int64())
}
