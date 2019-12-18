package day18

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	tests := []struct {
		maze  string
		steps int
	}{
		{maze: `#########
#b.A.@.a#
#########`, steps: 8},
		{maze: `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`, steps: 86},
		{maze: `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`, steps: 132},
		{maze: `#################
		#i.G..c...e..H.p#
		########.########
		#j.A..b...f..D.o#
		########@########
		#k.E..a...g..B.n#
		########.########
		#l.F..d...h..C.m#
		#################`, steps: 136},
		{maze: `########################
		#@..............ac.GI.b#
		###d#e#f################
		###A#B#C################
		###g#h#i################
		########################`, steps: 81},
	}

	for _, tt := range tests {
		t.Run(tt.maze, func(t *testing.T) {
			require.Equal(t, tt.steps, a(challenge.FromLiteral(tt.maze)))
		})
	}
}
