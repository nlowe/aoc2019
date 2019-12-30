package day22

import (
	"path/filepath"
	"testing"

	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/viper"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func TestFastFancyShuffleA(t *testing.T) {
	viper.Set("input", day22input())
	defer viper.Set("input", "")

	iterations = 1
	require.Equal(t, 3293, fastFancyShuffle(challenge.FromFile(), 10007, 2019))
}

func day22input() string {
	p, err := util.PkgPath(22)
	if err != nil {
		panic(err)
	}

	return filepath.Join(p, "input.txt")
}
