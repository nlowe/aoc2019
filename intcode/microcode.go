package intcode

import (
	"fmt"
	"time"

	"github.com/nlowe/aoc2019/intcode/instruction"
)

type opFunc func(m3 int, m2 int, m1 int, c *CPU) int

var microcode = map[int]opFunc{
	instruction.Add: func(m3, m2, m1 int, c *CPU) int {
		c.write(m3, 3, c.read(m1, 1)+c.read(m2, 2))
		return 1 + instruction.ArgCount(instruction.Add)
	},
	instruction.Mul: func(m3, m2, m1 int, c *CPU) int {
		c.write(m3, 3, c.read(m1, 1)*c.read(m2, 2))
		return 1 + instruction.ArgCount(instruction.Mul)
	},
	instruction.In: func(m3, m2, m1 int, c *CPU) int {
		select {
		case v := <-c.input:
			c.write(m1, 1, v)
			return 1 + instruction.ArgCount(instruction.In)
		case <-time.After(5 * time.Second):
			panic(fmt.Sprintf("no more input remaining after 5 seconds %s", c.debugState()))
		}
	},
	instruction.Out: func(m3, m2, m1 int, c *CPU) int {
		c.output <- c.read(m1, 1)
		return 1 + instruction.ArgCount(instruction.Out)
	},
	instruction.JT: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) != 0 {
			c.pc = c.read(m2, 2)
			return 0
		}

		return 1 + instruction.ArgCount(instruction.JT)
	},
	instruction.JF: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) == 0 {
			c.pc = c.read(m2, 2)
			return 0
		}

		return 1 + instruction.ArgCount(instruction.JF)
	},
	instruction.LT: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) < c.read(m2, 2) {
			c.write(m3, 3, 1)
		} else {
			c.write(m3, 3, 0)
		}

		return 1 + instruction.ArgCount(instruction.LT)
	},
	instruction.EQ: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) == c.read(m2, 2) {
			c.write(m3, 3, 1)
		} else {
			c.write(m3, 3, 0)
		}

		return 1 + instruction.ArgCount(instruction.EQ)
	},
	instruction.Rel: func(m3, m2, m1 int, c *CPU) int {
		c.relativeOffset += c.read(m1, 1)

		return 1 + instruction.ArgCount(instruction.Rel)
	},
	instruction.Halt: func(_, _, _ int, c *CPU) int {
		close(c.output)
		return 0
	},
}
