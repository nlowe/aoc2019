package day15

import "fmt"

const (
	moveNorth = iota + 1
	moveSouth
	moveWest
	moveEast
)

func backTrackingMove(move int) int {
	switch move {
	case moveNorth:
		return moveSouth
	case moveSouth:
		return moveNorth
	case moveWest:
		return moveEast
	case moveEast:
		return moveWest
	default:
		panic(fmt.Errorf("unknown move: %d", move))
	}
}

type robot struct {
	in  chan<- int
	out <-chan int

	X int
	Y int
}

func (r *robot) check(move int) int {
	r.in <- move
	status := <-r.out

	if status == statusOk || status == statusOxygenSystemFound {
		r.in <- backTrackingMove(move)
		<-r.out
	}

	return status
}

func (r *robot) move(move int) {
	switch move {
	case moveNorth:
		r.Y++
	case moveSouth:
		r.Y--
	case moveWest:
		r.X--
	case moveEast:
		r.X++
	}

	r.in <- move
	<-r.out
}
