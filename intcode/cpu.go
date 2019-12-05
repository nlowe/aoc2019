package intcode

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2019/util"
)

const (
	OpAdd  = 1
	OpMul  = 2
	OpIn   = 3
	OpOut  = 4
	OpHalt = 99

	ModeIndirect  = 0
	ModeImmediate = 1
)

type CPU struct {
	Memory []int

	inputs  <-chan int
	outputs chan<- int
	pc      int
}

func NewCPUForProgram(program string, inputs <-chan int) (*CPU, <-chan int) {
	parts := strings.Split(program, ",")
	memory := make([]int, len(parts))

	for i, op := range parts {
		memory[i] = util.MustAtoI(op)
	}

	output := make(chan int)
	return &CPU{Memory: memory, inputs: inputs, outputs: output}, output
}

func (c *CPU) Run() {
	for !c.IsHalted() {
		c.Step()
	}
}

func (c *CPU) Step() {
	m3, m2, m1, op := c.parseOp()
	switch op {
	case OpAdd:
		c.write(m3, c.Memory[c.pc+3], c.read(m1, c.Memory[c.pc+1])+c.read(m2, c.Memory[c.pc+2]))
		c.pc += 4
	case OpMul:
		c.write(m3, c.Memory[c.pc+3], c.read(m1, c.Memory[c.pc+1])*c.read(m2, c.Memory[c.pc+2]))
		c.pc += 4
	case OpIn:
		select {
		case v := <-c.inputs:
			c.write(m1, c.Memory[c.pc+1], v)
			c.pc += 2
		default:
			c.hcf("no more input remaining")
		}
	case OpOut:
		c.outputs <- c.read(m1, c.Memory[c.pc+1])
		c.pc += 2
	case OpHalt:
		close(c.outputs)
		return
	default:
		c.hcf("unknown opcode")
	}
}

func (c *CPU) parseOp() (int, int, int, int) {
	op := c.Memory[c.pc]

	return op / 10000, (op / 1000) % 10, (op / 100) % 10, op % 100
}

func (c *CPU) read(mode, target int) int {
	switch mode {
	case ModeIndirect:
		return c.Memory[target]
	case ModeImmediate:
		return target
	default:
		c.hcf(fmt.Sprintf("read: unknown mode: %d", mode))
		// c.hcf always panics, this is just to satisfy the compiler
		panic(nil)
	}
}

func (c *CPU) write(mode, target, value int) {
	switch mode {
	case ModeIndirect:
		c.Memory[target] = value
	case ModeImmediate:
		c.hcf(fmt.Sprintf("write: unsupported mode: Immediate"))
	default:
		c.hcf(fmt.Sprintf("write: unknown mode: %d", mode))
	}
}

func (c *CPU) hcf(msg string) {
	panic(fmt.Sprintf("%s [Op: %05d, PC: %d] (Memory: %+v)", msg, c.Memory[c.pc], c.pc, c.Memory))
}

func (c *CPU) IsHalted() bool {
	return c.Memory[c.pc] == OpHalt
}
