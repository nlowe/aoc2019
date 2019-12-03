package day2

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2019/util"
)

const (
	OpAdd  = 1
	OpMul  = 2
	OpHalt = 99
)

type CPU struct {
	Memory []int
	pc     int
}

func NewCPUForProgram(program string) *CPU {
	parts := strings.Split(program, ",")
	memory := make([]int, len(parts))

	for i, op := range parts {
		memory[i] = util.MustAtoI(op)
	}

	return &CPU{Memory: memory}
}

func (c *CPU) Run() {
	for !c.IsHalted() {
		c.Step()
	}
}

func (c *CPU) Step() {
	switch c.Memory[c.pc] {
	case OpAdd:
		c.Memory[c.Memory[c.pc+3]] = c.Memory[c.Memory[c.pc+1]] + c.Memory[c.Memory[c.pc+2]]
	case OpMul:
		c.Memory[c.Memory[c.pc+3]] = c.Memory[c.Memory[c.pc+1]] * c.Memory[c.Memory[c.pc+2]]
	case OpHalt:
		return
	default:
		panic(fmt.Errorf("unknown opcode: %d [PC: %d] (Memory: %+v)", c.Memory[c.pc], c.pc, c.Memory))
	}

	c.pc += 4
}

func (c *CPU) IsHalted() bool {
	return c.Memory[c.pc] == OpHalt
}
