package day18

import (
	"fmt"

	"github.com/nlowe/aoc2019/challenge"
)

const (
	tileOpen  = '.'
	tileWall  = '#'
	tileStart = '@'
)

type keyset int32

const NoKeys = keyset(0)

func keybit(k rune) int32 {
	return 1 << (k - 'a')
}

func (k *keyset) add(key rune) {
	if key < 'a' || key > 'z' {
		panic(fmt.Errorf("invalid key: %s", string(key)))
	}

	*k |= keyset(keybit(key))
}

func (k *keyset) remove(key rune) {
	if key < 'a' || key > 'z' {
		panic(fmt.Errorf("invalid key: %s", string(key)))
	}

	*k &^= keyset(keybit(key))
}

func (k keyset) has(key rune) bool {
	if key < 'a' || key > 'z' {
		panic(fmt.Errorf("invalid key: %s", string(key)))
	}

	bit := keybit(key)
	return int32(k)&bit == bit
}

func (k keyset) toRuneSlice() (result []rune) {
	for r := 'a'; r <= 'z'; r++ {
		if k.has(r) {
			result = append(result, r)
		}
	}

	return
}

type tile struct {
	x int
	y int
	r rune

	m *maze
}

func (t *tile) isKey() bool {
	return t.r >= 'a' && t.r <= 'z'
}

func (t *tile) isDoor() bool {
	return t.r >= 'A' && t.r <= 'Z'
}

type maze struct {
	m map[int]map[int]*tile

	keys map[rune]*tile

	startX int
	startY int
}

func ParseMaze(challenge *challenge.Input) *maze {
	result := &maze{m: map[int]map[int]*tile{}, keys: map[rune]*tile{}}
	y := 0
	for line := range challenge.Lines() {
		x := 0
		for _, r := range line {
			if _, ok := result.m[x]; !ok {
				result.m[x] = map[int]*tile{}
			}

			result.m[x][y] = &tile{
				x: x,
				y: y,
				r: r,
				m: result,
			}

			if r == tileStart {
				result.startX = x
				result.startY = y
			}

			if r >= 'a' && r <= 'z' {
				result.keys[r] = result.m[x][y]
			}

			x++
		}

		y++
	}

	return result
}

func (s *maze) tileAt(x, y int) *tile {
	if _, ok := s.m[x]; !ok {
		return nil
	}

	return s.m[x][y]
}
