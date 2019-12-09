package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

const (
	problemTemplate = `package day{{ .N }}

import (
    "fmt"

    "github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var {{ .AB }} = &cobra.Command{
    Use:   "{{ .N }}{{ .AB }}",
    Short: "Day {{ .N }}, Problem {{ .AB }}",
    Run: func(_ *cobra.Command, _ []string) {
        fmt.Printf("Answer: %d\n", {{ .AB | toLower }}(challenge.FromFile()))
    },
}

func {{ .AB | toLower }}(challenge *challenge.Input) int {
    return 0
}
`

	testTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/nlowe/aoc2019/challenge"
	"github.com/stretchr/testify/require"
)

func Test{{ .AB }}(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := {{ .AB | toLower }}(input)

	require.Equal(t, 42, result)
}
`
)

type metadata struct {
	N  int
	AB string
}

func main() {
	if len(os.Args) != 3 {
		abort(fmt.Errorf("expected 3 args but got %d", len(os.Args)))
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		abort(err)
	}

	ab := strings.ToUpper(os.Args[2])
	if !strings.ContainsAny(ab, "AB") {
		abort(fmt.Errorf("unknown problem segment: %s", ab))
	}

	p := pkgPath(n)
	if err := os.MkdirAll(p, 0744); err != nil {
		abort(err)
	}

	m := metadata{N: n, AB: ab}

	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}

	challengeTemplate := template.Must(template.New("challenge").Funcs(funcs).Parse(problemTemplate))
	cf, err := os.Create(filepath.Join(p, fmt.Sprintf("%s.go", strings.ToLower(ab))))
	if err != nil {
		abort(err)
	}

	defer mustClose(cf)
	if err := challengeTemplate.Execute(cf, m); err != nil {
		abort(err)
	}

	testTemplate := template.Must(template.New("test").Funcs(funcs).Parse(testTemplate))
	tf, err := os.Create(filepath.Join(p, fmt.Sprintf("%s_test.go", strings.ToLower(ab))))
	if err != nil {
		abort(err)
	}

	defer mustClose(tf)
	if err := testTemplate.Execute(tf, m); err != nil {
		abort(err)
	}

	goimports := exec.Command("goimports", "-w", p)
	if err := goimports.Run(); err != nil {
		abort(err)
	}

	input, err := os.OpenFile(filepath.Join(p, "input.txt"), os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		abort(err)
	}
	defer mustClose(input)

	fmt.Printf("Generated problem %s for day %d. Be sure to add it to main.go\n", ab, n)

	// TODO: Can we modify main.go easily?
}

func pkgPath(day int) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		abort(fmt.Errorf("failed to generate package path"))
	}

	return filepath.Join(filepath.Dir(filepath.Dir(filename)), "challenge", fmt.Sprintf("day%d", day))
}

func abort(err error) {
	fmt.Printf("%s\n\nsyntax: go run gen/problem.go <day> <a|b>\n", err.Error())
	os.Exit(1)
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		panic(fmt.Errorf("error closing io.Closer: %w", err))
	}
}
