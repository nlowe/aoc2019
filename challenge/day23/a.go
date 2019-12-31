package day23

import (
	"fmt"
	"sync"

	"github.com/nlowe/aoc2019/intcode"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "23a",
	Short: "Day 23, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
	},
}

const (
	networkSize = 50
)

func a(challenge *challenge.Input) int {
	program := <-challenge.Lines()
	computers := make([]*intcode.CPU, networkSize)
	inputs := make([]chan int, networkSize)
	outputs := make([]<-chan int, networkSize)

	startup := sync.WaitGroup{}
	startup.Add(networkSize)

	for i := 0; i < networkSize; i++ {
		inputs[i] = make(chan int)
		computers[i], outputs[i] = intcode.NewCPUForProgram(program, inputs[i])

		go func(i int) {
			// Boot the nic and assign the ID
			fmt.Printf("[%2d] booting...\n", i)
			go computers[i].Run()
			inputs[i] <- i
			startup.Done()
		}(i)
	}

	startup.Wait()

	return <-NewRouter(inputs, outputs).RouteTraffic()
}
