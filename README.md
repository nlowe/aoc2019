# Advent of Code 2019

[![](https://github.com/nlowe/aoc2019/workflows/CI/badge.svg)](https://github.com/nlowe/aoc2019/actions) [![Coverage Status](https://coveralls.io/repos/github/nlowe/aoc2019/badge.svg?branch=master)](https://coveralls.io/github/nlowe/aoc2019?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/nlowe/aoc2019)](https://goreportcard.com/report/github.com/nlowe/aoc2019) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

Solutions for the 2019 Advent of Code

## Building

This project makes use of Go 1.13.

```bash
go mod download
go test ./...
```

## Running the Solutions

To run a solution, use the problem name followed by the path to an input file.

For example, to run problem 2a:

```bash
$ go run ./main.go 2a ./day2/input.txt
Answer: 9633
Took 999.4µs
```

## Adding New Solutions

A generator program is included in `gen/problem.go` that makes templates for each day. For
example, `go run gen/problem.go 9 a` will generate the following files:

* `challenge/day9/a.go`: The main problem implementation, containing a cobra command `A` and the implementation `func a(*challenge.Input) int`
* `challenge/day9/a_test.go`: A basic test template
* `challenge/day9/input.txt`: The challenge input

I don't yet have a way to register the cobra command in `main.go` automatically.

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.
