package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/output"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "17b",
	Short: "Day 17, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
	},
}

const (
	rotateLeft  = "L"
	rotateRight = "R"
)

func b(challenge *challenge.Input) int {
	program := <-challenge.Lines()

	s, _, _ := parseScaffolding(program)
	p := makePath(s)

	main, a, b, c := compress(p)

	in := make(chan int)
	cpu, out := intcode.NewCPUForProgram(program, in)
	cpu.Memory[0] = 2

	result := 0
	wg := output.Each(out, func(v int) {
		result = v
	})

	go cpu.Run()

	send(main, in)
	send(a, in)
	send(b, in)
	send(c, in)
	send("n", in)

	wg.Wait()

	return result
}

func send(cmd string, ch chan<- int) {
	for _, v := range cmd {
		ch <- int(v)
	}

	ch <- tileNewline
}

type robot struct {
	x    int
	y    int
	face rune
}

func (r *robot) walk(s *scaffolding) int {
	dx := 0
	dy := 0
	switch r.face {
	case robotUp:
		dy = -1
	case robotDown:
		dy = 1
	case robotLeft:
		dx = -1
	case robotRight:
		dx = 1
	}

	count := 0
	for s.get(r.x+dx, r.y+dy) != nil {
		count++
		r.x += dx
		r.y += dy
	}

	return count
}

func (r *robot) rotateToNextPath(s *scaffolding) string {
	switch r.face {
	case robotUp:
		if s.get(r.x-1, r.y) != nil {
			r.face = robotLeft
			return rotateLeft
		}

		r.face = robotRight
		return rotateRight
	case robotRight:
		if s.get(r.x, r.y-1) != nil {
			r.face = robotUp
			return rotateLeft
		}

		r.face = robotDown
		return rotateRight
	case robotDown:
		if s.get(r.x+1, r.y) != nil {
			r.face = robotRight
			return rotateLeft
		}

		r.face = robotLeft
		return rotateRight
	case robotLeft:
		if s.get(r.x, r.y+1) != nil {
			r.face = robotDown
			return rotateLeft
		}

		r.face = robotUp
		return rotateRight
	default:
		panic(fmt.Errorf("unknown face: %s", string(r.face)))
	}
}

func (r *robot) atEnd(s *scaffolding) bool {
	return len(s.neighbors(r.x, r.y)) == 1
}

func makePath(s *scaffolding) string {

	r := &robot{
		x:    s.robotX,
		y:    s.robotY,
		face: s.robotFace,
	}

	instructions := []string{
		r.rotateToNextPath(s),
		strconv.Itoa(r.walk(s)),
	}

	for !r.atEnd(s) {
		instructions = append(instructions, r.rotateToNextPath(s))
		instructions = append(instructions, strconv.Itoa(r.walk(s)))
	}

	return strings.Join(instructions, ",")
}

type compressionNode struct {
	order int
	base  string
	left  *compressionNode
	right *compressionNode
}

func (c *compressionNode) String() string {
	if c.left == nil && c.right == nil {
		return c.base
	}

	return fmt.Sprintf("%s,%s", c.left.String(), c.right.String())
}

func compress(program string) (main string, a string, b string, c string) {
	// TODO: This is ripped from https://www.reddit.com/r/adventofcode/comments/ebr7dg/2019_day_17_solutions/fb7ymcw?utm_source=share&utm_medium=web2x
	//       At some point, I want to figure out a proper way to do these, since Go's regex library doesn't support
	//       backreferences which this expression uses (so we need to use a third party package)
	re := regexp2.MustCompile(`^(.{1,20})\1*(.{1,20})(?:\1|\2)*(.{1,20})(?:\1|\2|\3)*$`, 0)

	m, err := re.FindStringMatch(program + ",")
	if err != nil {
		panic(err)
	}

	if m == nil {
		panic("failed to extract functions")
	}

	functions := m.Groups()
	a = strings.TrimSuffix(functions[1].Captures[0].String(), ",")
	b = strings.TrimSuffix(functions[2].Captures[0].String(), ",")
	c = strings.TrimSuffix(functions[3].Captures[0].String(), ",")

	main = strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				program, a, "A",
			),
			b, "B",
		),
		c, "C",
	)

	return
}
