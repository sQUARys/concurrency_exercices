package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}

	}()

	// DONE: if there is no value on channel, do not block.
	for {

		select {
		case m := <-ch:
			fmt.Println(m)
			return
		default:
			fmt.Println("no messages")
		}
	}
}
