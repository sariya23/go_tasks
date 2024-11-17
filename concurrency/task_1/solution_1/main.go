package main

import (
	"fmt"
)

func calculations(v int) <-chan int {
	ch := make(chan int)
	go func() {
		ch <- v * v
		close(ch)
	}()

	return ch
}

func funOut(channels ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for _, ch := range channels {
			for el := range ch {
				out <- el
			}
		}
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
	for v := range funOut(calcChannels...) {
		sum += v
	}

	fmt.Printf("result: %d\n", sum)
}
