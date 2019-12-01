# Advent of Code 2019

![](https://github.com/nlowe/aoc2019/workflows/CI/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/nlowe/aoc2019)](https://goreportcard.com/report/github.com/nlowe/aoc2019) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

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

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.
