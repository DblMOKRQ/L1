package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type incCounter struct {
	value int32
}

func (c *incCounter) Inc() {
	atomic.AddInt32(&c.value, 1)
}
func (c *incCounter) Get() int32 {
	return atomic.LoadInt32(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := &incCounter{}

	numGoroutines := 10
	incrementsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Итоговое значение счётчика: %d\n", counter.Get())

}
