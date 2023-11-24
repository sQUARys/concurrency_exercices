package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// DONE: implement timeout for recv on channel ch

loop:
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			break loop
		default:
			fmt.Println("timout start")
			time.Sleep(time.Second)
			fmt.Println("timout end")
		}
	}
}
