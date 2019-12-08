package main

import (
	"fmt"
	"time"

	"github.com/nlowe/aoc2019/challenge/day7"

	"github.com/nlowe/aoc2019/challenge/day1"
	"github.com/nlowe/aoc2019/challenge/day2"
	"github.com/nlowe/aoc2019/challenge/day3"
	"github.com/nlowe/aoc2019/challenge/day4"
	"github.com/nlowe/aoc2019/challenge/day5"
	"github.com/nlowe/aoc2019/challenge/day6"
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
