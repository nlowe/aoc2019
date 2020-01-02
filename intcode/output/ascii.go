package output

import "strings"

func ToASCII(out <-chan int) <-chan string {
	result := make(chan string)

	go func() {
		buf := strings.Builder{}
		for r := range out {
			if r == '\n' {
				result <- buf.String()
				buf.Reset()
			} else {
				buf.WriteRune(rune(r))
			}
		}

		if buf.Len() > 0 {
			result <- buf.String()
		}

		close(result)
	}()

	return result
}
