package input

func NewFixed(inputs ...int) <-chan int {
	result := make(chan int, len(inputs))
	for _, v := range inputs {
		result <- v
	}

	return result
}
