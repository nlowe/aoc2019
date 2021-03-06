package main

import (
	"fmt"
	"time"

	"github.com/nlowe/aoc2019/challenge/day1"
	"github.com/nlowe/aoc2019/challenge/day10"
	"github.com/nlowe/aoc2019/challenge/day11"
	"github.com/nlowe/aoc2019/challenge/day12"
	"github.com/nlowe/aoc2019/challenge/day13"
	"github.com/nlowe/aoc2019/challenge/day14"
	"github.com/nlowe/aoc2019/challenge/day15"
	"github.com/nlowe/aoc2019/challenge/day16"
	"github.com/nlowe/aoc2019/challenge/day17"
	"github.com/nlowe/aoc2019/challenge/day18"
	"github.com/nlowe/aoc2019/challenge/day19"
	"github.com/nlowe/aoc2019/challenge/day2"
	"github.com/nlowe/aoc2019/challenge/day20"
	"github.com/nlowe/aoc2019/challenge/day21"
	"github.com/nlowe/aoc2019/challenge/day22"
	"github.com/nlowe/aoc2019/challenge/day23"
	"github.com/nlowe/aoc2019/challenge/day24"
	"github.com/nlowe/aoc2019/challenge/day25"
	"github.com/nlowe/aoc2019/challenge/day3"
	"github.com/nlowe/aoc2019/challenge/day4"
	"github.com/nlowe/aoc2019/challenge/day5"
	"github.com/nlowe/aoc2019/challenge/day6"
	"github.com/nlowe/aoc2019/challenge/day7"
	"github.com/nlowe/aoc2019/challenge/day8"
	"github.com/nlowe/aoc2019/challenge/day9"
	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type prof interface {
	Stop()
}

var (
	start    time.Time
	profiler prof

	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 Solutions",
		Long:  "Golang implementations for the 2019 Advent of Code problems",
		Args:  cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if viper.GetBool("profile") {
				profiler = profile.Start()
			}

			start = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			if profiler != nil {
				profiler.Stop()
			}

			fmt.Printf("Took %s\n", time.Since(start))
		},
	}
)

func init() {
	rootCmd.AddCommand(
		day1.A, day1.B,
		day2.A, day2.B,
		day3.A, day3.B,
		day4.A, day4.B,
		day5.A, day5.B,
		day6.A, day6.B,
		day7.A, day7.B,
		day8.A, day8.B,
		day9.A, day9.B,
		day10.A, day10.B,
		day11.A, day11.B,
		day12.A, day12.B,
		day13.A, day13.B,
		day14.A, day14.B,
		day15.A, day15.B,
		day16.A, day16.B,
		day17.A, day17.B,
		day18.A, day18.B,
		day19.A, day19.B,
		day20.A, day20.B,
		day21.A, day21.B,
		day22.A, day22.B,
		day23.A, day23.B,
		day24.A, day24.B,
		day25.A,
	)

	flags := rootCmd.PersistentFlags()
	flags.StringP("input", "i", "", "Input File to read")
	if err := rootCmd.MarkPersistentFlagRequired("input"); err != nil {
		panic(err)
	}

	flags.Bool("profile", false, "Profile implementation performance")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
