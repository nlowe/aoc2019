package output

import (
	"sync"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Each(outputs <-chan int, cb func(int)) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	go func() {
		for code := range outputs {
			cb(code)
		}

		wg.Done()
	}()

	return wg
}

func Expect(t *testing.T, outputs <-chan int, expected ...int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for v := range outputs {
			assert.Equal(t, expected[0], v)
		}
		wg.Done()
	}()

	return wg
}

func Single(outputs <-chan int, result *int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		*result = <-outputs
		wg.Done()
	}()

	return wg
}
