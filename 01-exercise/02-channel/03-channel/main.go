package main

import (
	"fmt"
)

func main() {
	const countOfIterations = 6

	ch := make(chan int, countOfIterations)

	go func() {
		defer close(ch)

		// DONE: send all iterator values on channel without blocking
		for i := 0; i < countOfIterations; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
