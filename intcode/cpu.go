package intcode

import (
	"fmt"
	"strings"
	"time"

	"github.com/nlowe/aoc2019/util"
)

const (
	OpAdd  = 1
	OpMul  = 2
	OpIn   = 3
	OpOut  = 4
	OpJT   = 5
	OpJF   = 6
	OpLT   = 7
	OpEQ   = 8
	OpHalt = 99

	ModeIndirect  = 0
	ModeImmediate = 1
)

type CPU struct {
	Memory []int

	input  <-chan int
	output chan<- int
	pc     int
}

func NewCPUForProgram(program string, inputs <-chan int) (*CPU, <-chan int) {
	parts := strings.Split(program, ",")
	memory := make([]int, len(parts))

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
	switch op {
	case OpAdd:
		c.write(m3, 3, c.read(m1, 1)+c.read(m2, 2))
		c.pc += 4
	case OpMul:
		c.write(m3, 3, c.read(m1, 1)*c.read(m2, 2))
		c.pc += 4
	case OpIn:
		select {
		case v := <-c.input:
			c.write(m1, 1, v)
			c.pc += 2
		case <-time.After(5 * time.Second):
			panic(fmt.Sprintf("no more input remaining after 5 seconds %s", c.debugState()))
		}
	case OpOut:
		c.output <- c.read(m1, 1)
		c.pc += 2
	case OpJT:
		if c.read(m1, 1) != 0 {
			c.pc = c.read(m2, 2)
		} else {
			c.pc += 3
		}
	case OpJF:
		if c.read(m1, 1) == 0 {
			c.pc = c.read(m2, 2)
		} else {
			c.pc += 3
		}
	case OpLT:
		if c.read(m1, 1) < c.read(m2, 2) {
			c.write(m3, 3, 1)
		} else {
			c.write(m3, 3, 0)
		}

		c.pc += 4
	case OpEQ:
		if c.read(m1, 1) == c.read(m2, 2) {
			c.write(m3, 3, 1)
		} else {
			c.write(m3, 3, 0)
		}

		c.pc += 4
	case OpHalt:
		close(c.output)
		return
	default:
		panic(fmt.Sprintf("unknown opcode %s", c.debugState()))
	}
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
