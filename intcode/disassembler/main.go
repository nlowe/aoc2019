package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nlowe/aoc2019/intcode/instruction"

	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/util"
)

const symData = ".db"

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

		m3, m2, m1, op := instruction.Parse(instr)
		name, found := instruction.NameOf(op)
		if !found {
			program.WriteString(symData)
			program.WriteString(fmt.Sprintf(" $%d\n", op))
			continue
		}

		program.WriteString(name)
		modes := []int{m1, m2, m3}
		var args []string
		for arg := 0; arg < instruction.ArgCount(op); arg++ {
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
