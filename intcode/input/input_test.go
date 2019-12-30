package input

import (
	"testing"

	"github.com/nlowe/aoc2019/intcode/output"

	"github.com/stretchr/testify/require"
)

func TestPrefix(t *testing.T) {
	other := make(chan int)
	go func() {
		other <- 42
	}()

	sut := Prefix(24, other)

	require.Equal(t, 24, <-sut)
	require.Equal(t, 42, <-sut)
}

func TestASCIIWrapper(t *testing.T) {
	w := make(chan string)
	sut := ASCIIWrapper(w)

	go func() {
		w <- "hello"
		w <- "world"
	}()

	output.Expect(t, sut, 'h', 'e', 'l', 'l', 'o', '\n', 'w', 'o', 'r', 'l', 'd', '\n')
}
