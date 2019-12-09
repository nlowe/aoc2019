package intcode

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2019/util"
)

const (
	ModeIndirect  = 0
	ModeImmediate = 1
	ModeRelative  = 2

	RAMSize = 8192
)

type CPU struct {
	Memory []int

	input  <-chan int
	output chan<- int

	pc             int
	relativeOffset int
}

func NewCPUForProgram(program string, inputs <-chan int) (*CPU, <-chan int) {
	parts := strings.Split(program, ",")
	memory := make([]int, RAMSize)

	for i, op := range parts {
		memory[i] = util.MustAtoI(op)
	}

	output := make(chan int)
	return &CPU{Memory: memory, input: inputs, output: output}, output
}

func (c *CPU) Run() {
	for !c.IsHalted() {
		c.Step()
	}

	close(c.output)
}

func (c *CPU) Step() {
	m3, m2, m1, op := c.parseOp()
	impl, ok := opTable[op]
	if !ok {
		panic(fmt.Sprintf("unknown opcode %s", c.debugState()))
	}

	c.pc += impl(m3, m2, m1, c)
}

func (c *CPU) parseOp() (int, int, int, int) {
	op := c.Memory[c.pc]

	return op / 10000, (op / 1000) % 10, (op / 100) % 10, op % 100
}

func (c *CPU) read(mode, offset int) int {
	switch mode {
	case ModeIndirect:
		return c.Memory[c.Memory[c.pc+offset]]
	case ModeImmediate:
		return c.Memory[c.pc+offset]
	case ModeRelative:
		return c.Memory[c.relativeOffset+c.Memory[c.pc+offset]]
	default:
		panic(fmt.Sprintf("read: unknown mode: %d %s", mode, c.debugState()))
	}
}

func (c *CPU) write(mode, offset, value int) {
	switch mode {
	case ModeIndirect:
		c.Memory[c.Memory[c.pc+offset]] = value
	case ModeImmediate:
		panic(fmt.Sprintf("write: unsupported mode: Immediate %s", c.debugState()))
	case ModeRelative:
		c.Memory[c.relativeOffset+c.Memory[c.pc+offset]] = value
	default:
		panic(fmt.Sprintf("write: unknown mode: %d %s", mode, c.debugState()))
	}
}

func (c *CPU) debugState() string {
	return fmt.Sprintf("[Op: %05d, PC: %d] (Memory: %+v)", c.Memory[c.pc], c.pc, c.Memory)
}

func (c *CPU) IsHalted() bool {
	return c.Memory[c.pc] == OpHalt
}
