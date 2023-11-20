package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		valueCh = make(chan int)
		wg      sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			valueCh <- i
		}

		close(valueCh)
	}()

	for v := range valueCh {
		fmt.Println(v)
	}

	wg.Wait()

	// DONE: range over channel to recv values

}
