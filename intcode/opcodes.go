package intcode

import (
	"fmt"
	"time"
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
	OpRel  = 9
	OpHalt = 99
)

type opFunc func(m3 int, m2 int, m1 int, c *CPU) int

var opTable = map[int]opFunc{
	OpAdd: func(m3, m2, m1 int, c *CPU) int {
		c.write(m3, 3, c.read(m1, 1)+c.read(m2, 2))
		return 4
	},
	OpMul: func(m3, m2, m1 int, c *CPU) int {
		c.write(m3, 3, c.read(m1, 1)*c.read(m2, 2))
		return 4
	},
	OpIn: func(m3, m2, m1 int, c *CPU) int {
		select {
		case v := <-c.input:
			c.write(m1, 1, v)
			return 2
		case <-time.After(5 * time.Second):
			panic(fmt.Sprintf("no more input remaining after 5 seconds %s", c.debugState()))
		}
	},
	OpOut: func(m3, m2, m1 int, c *CPU) int {
		c.output <- c.read(m1, 1)
		return 2
	},
	OpJT: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) != 0 {
			c.pc = c.read(m2, 2)
			return 0
		}

		return 3
	},
	OpJF: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) == 0 {
			c.pc = c.read(m2, 2)
			return 0
		}

		return 3
	},
	OpLT: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) < c.read(m2, 2) {
			c.write(m3, 3, 1)
		} else {
			c.write(m3, 3, 0)
		}

		return 4
	},
	OpEQ: func(m3, m2, m1 int, c *CPU) int {
		if c.read(m1, 1) == c.read(m2, 2) {
			c.write(m3, 3, 1)
		} else {
			c.write(m3, 3, 0)
		}

		return 4
	},
	OpRel: func(m3, m2, m1 int, c *CPU) int {
		c.relativeOffset += c.read(m1, 1)

		return 2
	},
	OpHalt: func(_, _, _ int, c *CPU) int {
		close(c.output)
		return 0
	},
}
