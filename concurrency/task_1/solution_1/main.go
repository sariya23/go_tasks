package main

import (
	"fmt"
	"sync"
)

func calculations(v int) <-chan int {
	ch := make(chan int)
	go func() {
		ch <- v * v
		close(ch)
	}()

	return ch
}

func funIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for el := range ch {
				out <- el
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	calcChannels := make([]<-chan int, 0, 3)
	for i := 0; i < 3; i++ {
		calcChannels = append(calcChannels, calculations(i))
	}

	var sum int
	for v := range funIn(calcChannels...) {
		sum += v
	}

	fmt.Printf("result: %d\n", sum)
}
