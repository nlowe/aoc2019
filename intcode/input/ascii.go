package input

func ASCIIWrapper(in <-chan string) <-chan int {
	wrapped := make(chan int)

	go func() {
		for line := range in {
			for _, r := range line {
				wrapped <- int(r)
			}
			wrapped <- '\n'
		}
	}()

	return wrapped
}
