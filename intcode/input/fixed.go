package input

func NewFixed(inputs ...int) <-chan int {
	result := make(chan int, len(inputs))
	for _, v := range inputs {
		result <- v
	}

	return result
}

func Prefix(prefix int, rest <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		result <- prefix
		for v := range rest {
			result <- v
		}

		close(result)
	}()

	return result
}
