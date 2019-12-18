package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	main, a, b, c := compress("L,8,R,12,R,12,R,10,R,10,R,12,R,10,L,8,R,12,R,12,R,10,R,10,R,12,R,10,L,10,R,10,L,6,L,10,R,10,L,6,R,10,R,12,R,10,L,8,R,12,R,12,R,10,R,10,R,12,R,10,L,10,R,10,L,6")

	assert.Equal(t, "A,B,A,B,C,C,B,A,B,C", main, "main")
	assert.Equal(t, "L,8,R,12,R,12,R,10", a, "function A")
	assert.Equal(t, "R,10,R,12,R,10", b, "function B")
	assert.Equal(t, "L,10,R,10,L,6", c, "function C")
}
