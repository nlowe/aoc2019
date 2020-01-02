package day25

import (
	"fmt"
	"strings"
	"time"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/nlowe/aoc2019/intcode"
	"github.com/nlowe/aoc2019/intcode/input"
	"github.com/nlowe/aoc2019/intcode/output"
	"github.com/nlowe/aoc2019/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var A = &cobra.Command{
	Use:   "25a",
	Short: "Day 25, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

const promptCommand = "Command?"

func init() {
	flags := A.Flags()

	flags.BoolP("verbose", "v", false, "Show output from the game")
	flags.Bool("explore-only", false, "Only explore the ship, don't take items")

	_ = viper.BindPFlags(flags)
}

// Not all items are safe to pick up. Some abort the game immediately, others
// put the robot into an unrecoverable state. These are the ones I found for
// my puzzle, yours may be different. Of all inputs I've seen, you should have
// 8 items left over.
var itemBlacklist = []string{
	"infinite loop",
	"giant electromagnet",
	"escape pod",
	"molten lava",
	"photons",
}

type robot struct {
	foundItems               []string
	pathToSecurityCheckpoint []string
	foundSecurityCheckpoint  bool
}

func a(challenge *challenge.Input) int {
	in := make(chan string)
	cpu, rawOut := intcode.NewCPUForProgram(<-challenge.Lines(), input.ASCIIWrapper(in))
	cpu.WatchdogTimeout = 1 * time.Hour

	out := output.ToASCII(rawOut)

	go cpu.Run()

	r := &robot{}
	_, directions, items := r.observe(out)
	r.takeSafeItems(in, out, items)
	for _, d := range directions {
		r.explore(in, out, d)
	}

	fmt.Println("-=-=- Ship Explored -=-=-")
	fmt.Printf("Found: [%s]\n", strings.Join(r.foundItems, ", "))
	fmt.Printf("Path to security checkpoint: %+v\n", r.pathToSecurityCheckpoint)

	fmt.Println("-=-=- Navigating to Security Checkpoint -=-=-")
	for _, d := range r.pathToSecurityCheckpoint {
		in <- d
		_, directions, _ = r.observe(out)
	}

	last := r.pathToSecurityCheckpoint[len(r.pathToSecurityCheckpoint)-1]
	move := "unknown"
	for _, d := range directions {
		if d == backTrackingMove(last) {
			continue
		}

		move = d
		break
	}

	fmt.Println("Security Check should be triggered by going", move)
	in <- move
	advanceToCommandPrompt(out)

	return r.bruteForceKey(in, out, move)
}

func (r *robot) bruteForceKey(in chan<- string, out <-chan string, checkMove string) int {
	fmt.Printf("Brute-forcing key, 2^%d==%d combinations to try\n", len(r.foundItems), 1<<len(r.foundItems))
	fmt.Println("Dropping all items")
	for _, item := range r.foundItems {
		in <- fmt.Sprintf("drop %s", item)
		advanceToCommandPrompt(out)
	}

	for mask := 0; mask < 1<<len(r.foundItems); mask++ {
		try := r.pick(mask)
		fmt.Printf("Trying [%s]...\n", strings.Join(try, ","))
		for _, item := range try {
			in <- fmt.Sprintf("take %s", item)
			advanceToCommandPrompt(out)
		}

		in <- checkMove
		potentialAnswer, correct := tryReadAnswer(out)
		if correct {
			fmt.Printf("Correct combination was [%s]\n", strings.Join(try, ","))
			return potentialAnswer
		}

		for _, item := range try {
			in <- fmt.Sprintf("drop %s", item)
			advanceToCommandPrompt(out)
		}
	}

	panic("no solution")
}

func (r *robot) pick(mask int) (result []string) {
	idx := 0
	for mask > 0 {
		var sel int
		sel, mask = mask&0b1, mask>>1
		if sel == 0b1 {
			result = append(result, r.foundItems[idx])
		}

		idx++
	}

	return
}

func (r *robot) observe(out <-chan string) (room string, directions []string, items []string) {
	line := ""
	room = "Unknown"
	findingDoors := false

	for line != promptCommand {
		var more bool
		line, more = <-out
		if !more {
			panic("EOF")
		}

		if viper.GetBool("verbose") {
			fmt.Printf("|\t%s\n", line)
		}
		if strings.HasPrefix(line, "== ") {
			room = strings.TrimSuffix(strings.TrimPrefix(line, "== "), " ==")
			fmt.Println("Found room", room)
		} else if strings.EqualFold("Doors here lead:", line) {
			findingDoors = true
		} else if strings.EqualFold("Items here:", line) {
			findingDoors = false
		} else if strings.HasPrefix(line, "- ") {
			itemOrDoor := strings.TrimPrefix(line, "- ")
			if findingDoors {
				directions = append(directions, itemOrDoor)
			} else {
				items = append(items, itemOrDoor)
			}
		}
	}

	return
}

func (r *robot) explore(in chan<- string, out <-chan string, direction string) {
	fmt.Println("Exploring", direction)
	in <- direction
	backwards := backTrackingMove(direction)

	if !r.foundSecurityCheckpoint {
		r.pathToSecurityCheckpoint = append(r.pathToSecurityCheckpoint, direction)
	}

	room, directions, items := r.observe(out)
	if room == "Security Checkpoint" {
		r.foundSecurityCheckpoint = true
	}

	for _, i := range items {
		fmt.Println("***Found", i, "in", room)
	}
	r.takeSafeItems(in, out, items)

	// For now, don't bother exploring past the security checkpoint since we
	// probably don't have all of the items just yet
	if room == "Security Checkpoint" {
		r.foundSecurityCheckpoint = true
	} else {
		for _, d := range directions {
			if d == backwards {
				continue
			}

			r.explore(in, out, d)
		}
	}

	fmt.Println("Back-tracking", backwards)
	if !r.foundSecurityCheckpoint {
		r.pathToSecurityCheckpoint = r.pathToSecurityCheckpoint[:len(r.pathToSecurityCheckpoint)-1]
	}
	in <- backwards
	advanceToCommandPrompt(out)
}

func backTrackingMove(direction string) string {
	switch direction {
	case "north":
		return "south"
	case "south":
		return "north"
	case "west":
		return "east"
	case "east":
		return "west"
	default:
		panic(fmt.Errorf("unknown direction: %s", direction))
	}
}

func advanceToCommandPrompt(out <-chan string) {
	line := ""
	for line != promptCommand {
		var more bool
		line, more = <-out

		if !more {
			panic("EOF")
		}

		if viper.GetBool("verbose") {
			fmt.Printf("|\t%s\n", line)
		}
	}
}

func tryReadAnswer(out <-chan string) (int, bool) {
	line := ""
	for line != promptCommand {
		var more bool
		line, more = <-out
		if !more {
			panic("EOF")
		}

		if viper.GetBool("verbose") {
			fmt.Printf("|\t%s\n", line)
		}

		if strings.HasPrefix(line, `"Oh, hello! You should be able to get in by typing `) {
			return util.MustAtoI(strings.TrimSuffix(strings.TrimPrefix(line, `"Oh, hello! You should be able to get in by typing `), ` on the keypad at the main airlock."`)), true
		}
	}

	return 0, false
}

func (r *robot) takeSafeItems(in chan<- string, out <-chan string, items []string) {
search:
	for _, item := range items {
		for _, banned := range itemBlacklist {
			if item == banned {
				fmt.Println(item, "is not safe to take, skipping...")
				continue search
			}
		}

		if !viper.GetBool("explore-only") {
			fmt.Println("Taking", item)
			r.foundItems = append(r.foundItems, item)
			in <- fmt.Sprintf("take %s", item)
			advanceToCommandPrompt(out)
		} else {
			fmt.Println("Would take", item)
		}
	}
}
