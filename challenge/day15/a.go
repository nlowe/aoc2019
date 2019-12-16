package day15

import (
	"fmt"

	"github.com/beefsack/go-astar"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "15a",
	Short: "Day 15, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

func a(challenge *challenge.Input) int {
	s, ox, oy := mapShip(challenge)

	// TODO: Figure out this weird off-by-one error
	_, distance, found := astar.Path(s.tileAt(0, 0), s.tileAt(ox, oy-1))

	if !found {
		panic(fmt.Errorf("no solution: could not find path to oxygen system"))
	}

	return 1 + int(distance)
}

func mapShip(challenge *challenge.Input) (*ship, int, int) {
	in := make(chan int)
	cpu, out := intcode.NewCPUForProgram(<-challenge.Lines(), in)
	defer cpu.Halt()
	go cpu.Run()

	r := &robot{
		in:  in,
		out: out,
	}

	ox := 0
	oy := 0
	s := &ship{m: map[int]map[int]*tile{}}
	s.set(0, 0, statusOk)

	explore(moveNorth, r, s, &ox, &oy)
	explore(moveEast, r, s, &ox, &oy)
	explore(moveSouth, r, s, &ox, &oy)
	explore(moveWest, r, s, &ox, &oy)

	return s, ox, oy
}

func explore(direction int, r *robot, s *ship, ox, oy *int) {
	tile := r.check(direction)
	switch direction {
	case moveNorth:
		s.set(r.X, r.Y+1, tile)
	case moveEast:
		s.set(r.X+1, r.Y, tile)
	case moveSouth:
		s.set(r.X, r.Y-1, tile)
	case moveWest:
		s.set(r.X-1, r.Y, tile)
	}

	if tile == statusWall {
		return
	}

	r.move(direction)
	if tile == statusOxygenSystemFound {
		*ox = r.X
		*oy = r.Y
	}

	back := backTrackingMove(direction)
	for m := moveNorth; m <= moveEast; m++ {
		if m == back {
			continue
		}

		explore(m, r, s, ox, oy)
	}
	r.move(back)
}
