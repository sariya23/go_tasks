package main

import (
	"fmt"
	"sync"
)

func calculations() <-chan int {
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	var sum int
	for v := range calculations() {
		sum += v
	}

	fmt.Printf("result: %d\n", sum)
}
