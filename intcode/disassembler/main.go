package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/util"
)

const symData = ".db"

var opcodeNames = map[int]string{
	intcode.OpAdd:  "ADD",
	intcode.OpMul:  "MUL",
	intcode.OpIn:   "IN ",
	intcode.OpOut:  "OUT",
	intcode.OpJT:   "JT ",
	intcode.OpJF:   "JF ",
	intcode.OpLT:   "LT ",
	intcode.OpEQ:   "EQ ",
	intcode.OpRel:  "REL",
	intcode.OpHalt: "HLT",
}

var opcodeArgs = map[int]int{
	intcode.OpAdd:  3,
	intcode.OpMul:  3,
	intcode.OpIn:   1,
	intcode.OpOut:  1,
	intcode.OpJT:   2,
	intcode.OpJF:   2,
	intcode.OpLT:   3,
	intcode.OpEQ:   3,
	intcode.OpRel:  1,
	intcode.OpHalt: 0,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("syntax: go run intcode/disassembler/main.go path/to/file.txt")
		os.Exit(1)
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(f), ",")
	memory := make([]int, len(parts))

	for i, op := range parts {
		memory[i] = util.MustAtoI(strings.TrimSpace(op))
	}

	program := strings.Builder{}
	pc := 0
	for pc < len(memory) {
		instr := memory[pc]
		pc++

		m3, m2, m1, op := instr/10000, (instr/1000)%10, (instr/100)%10, instr%100
		name, found := opcodeNames[op]
		if !found {
			program.WriteString(symData)
			program.WriteString(fmt.Sprintf(" $%d\n", op))
			continue
		}

		program.WriteString(name)
		modes := []int{m1, m2, m3}
		var args []string
		for arg := 0; arg < opcodeArgs[op]; arg++ {
			m := modes[arg]
			args = append(args, decorate(m, memory[pc]))
			pc++
		}

		program.WriteString(" ")
		program.WriteString(strings.Join(args, ", "))
		program.WriteString("\n")
	}

	fmt.Println(program.String())
}

func decorate(mode, offset int) string {
	switch mode {
	case intcode.ModeIndirect:
		return fmt.Sprintf("[%d]", offset)
	case intcode.ModeImmediate:
		return fmt.Sprintf("$%d", offset)
	case intcode.ModeRelative:
		return fmt.Sprintf("[rp+$%d]", offset)
	default:
		panic(fmt.Errorf("unknown mode: %d", mode))
	}
}
