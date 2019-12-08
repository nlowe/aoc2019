package input

import (
	"testing"

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
