package instruction

import "fmt"

const (
	Add = iota + 1
	Mul
	In
	Out
	JT
	JF
	LT
	EQ
	Rel

	Halt = 99
)

var (
	names = map[int]string{
		Add:  "ADD",
		Mul:  "MUL",
		In:   "IN ",
		Out:  "OUT",
		JT:   "JT ",
		JF:   "JF ",
		LT:   "LT ",
		EQ:   "EQ ",
		Rel:  "REL",
		Halt: "HLT",
	}

	argCounts = map[int]int{
		Add:  3,
		Mul:  3,
		In:   1,
		Out:  1,
		JT:   2,
		JF:   2,
		LT:   3,
		EQ:   3,
		Rel:  1,
		Halt: 0,
	}
)

func NameOf(instr int) (string, bool) {
	v, ok := names[instr]
	return v, ok
}

func ArgCount(instr int) int {
	v, ok := argCounts[instr]
	if !ok {
		panic(fmt.Errorf("unknown instruction: %d", instr))
	}

	return v
}
