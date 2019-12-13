package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/zellyn/kooky"
)

const (
	problemTemplate = `package day{{ .N }}

import (
    "fmt"

    "github.com/nlowe/aoc2019/challenge"
	"github.com/spf13/cobra"
)

var {{ .AB }} = &cobra.Command{
    Use:   "{{ .N }}{{ .AB | toLower }}",
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

	genFile(filepath.Join(p, fmt.Sprintf("%s.go", strings.ToLower(ab))), problemTemplate, funcs, m)
	genFile(filepath.Join(p, fmt.Sprintf("%s_test.go", strings.ToLower(ab))), testTemplate, funcs, m)

	goimports := exec.Command("goimports", "-w", p)
	if err := goimports.Run(); err != nil {
		abort(err)
	}

	if _, stat := os.Stat(filepath.Join(p, "input.txt")); os.IsNotExist(stat) {
		fmt.Println("fetching input for day", n)
		problemInput, err := getInput(n)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(filepath.Join(p, "input.txt"), problemInput, 0644); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("input already downloaded, skipping...")
	}

	fmt.Printf("Generated problem %s for day %d. Be sure to add it to main.go\n", ab, n)

	// TODO: Can we modify main.go easily?
}

func genFile(path, t string, funcs template.FuncMap, m metadata) {
	if _, stat := os.Stat(path); os.IsNotExist(stat) {
		fmt.Println("creating", path)
		t := template.Must(template.New(path).Funcs(funcs).Parse(t))
		cf, err := os.Create(path)
		if err != nil {
			abort(err)
		}

		defer mustClose(cf)
		if err := t.Execute(cf, m); err != nil {
			abort(err)
		}
	} else {
		fmt.Println(path, "already exists, skipping...")
	}
}

func chromeCookiePath() (string, error) {
	if p, set := os.LookupEnv("CHROME_PROFILE_PATH"); set {
		return filepath.Join(p, "Cookies"), nil
	}

	if runtime.GOOS == "windows" {
		localAppData, err := os.UserCacheDir()
		return filepath.Join(localAppData, "Google", "Chrome", "User Data", "Default", "Cookies"), err
	}

	return "", fmt.Errorf("chrome cookie path for GOOS %s not implemented, set CHROME_PROFILE_PATH instead", runtime.GOOS)
}

func getInput(day int) ([]byte, error) {
	cookiePath, err := chromeCookiePath()
	if err != nil {
		return nil, err
	}

	_, _ = os.UserConfigDir()
	_, _ = os.UserCacheDir()
	cookies, err := kooky.ReadChromeCookies(cookiePath, ".adventofcode.com", "session", time.Time{})
	if err != nil {
		return nil, err
	}

	if len(cookies) == 0 {
		return nil, fmt.Errorf("session cookie not found, ensure that you are logged in")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2019/day/%d/input", day), nil)
	if err != nil {
		return nil, err
	}

	sessionToken := cookies[0].HttpCookie()
	req.AddCookie(&sessionToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer mustClose(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		panic(fmt.Errorf("error closing io.Closer: %w", err))
	}
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
