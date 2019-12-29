package intcode

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nlowe/aoc2019/intcode/instruction"
	"github.com/nlowe/aoc2019/util"
)

const (
	ModeIndirect  = 0
	ModeImmediate = 1
	ModeRelative  = 2

	RAMSize = 8192
)

type CPU struct {
	Memory map[int]int

	input  <-chan int
	output chan<- int

	pc             int
	relativeOffset int

	ctx             context.Context
	halt            context.CancelFunc
	halted          bool
	WatchdogTimeout time.Duration
}

func NewCPUForProgram(program string, inputs <-chan int) (*CPU, <-chan int) {
	parts := strings.Split(program, ",")
	memory := map[int]int{}

	for i, op := range parts {
		memory[i] = util.MustAtoI(op)
	}

	ctx, cancel := context.WithCancel(context.Background())

	output := make(chan int)
	return &CPU{
		Memory:          memory,
		input:           inputs,
		output:          output,
		ctx:             ctx,
		halt:            cancel,
		WatchdogTimeout: 5 * time.Second,
	}, output
}

func (c *CPU) Halt() {
	c.halted = true
	c.halt()
}

func (c *CPU) Run() {
	for !c.IsHalted() {
		c.Step()
	}

	close(c.output)
}

func (c *CPU) Step() {
	m3, m2, m1, op := instruction.Parse(c.Memory[c.pc])
	impl, ok := microcode[op]
	if !ok {
		panic(fmt.Sprintf("unknown opcode %s", c.debugState()))
	}

	c.pc += impl(m3, m2, m1, c)
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
	return c.halted || c.Memory[c.pc] == instruction.Halt
}
