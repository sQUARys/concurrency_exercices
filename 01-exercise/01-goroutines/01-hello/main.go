package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// DONE: write goroutine with different variants for function call.

	go fun("basic goroutine call")

	go func(s string) {
		fun(s)
	}("lambda func goroutine call")

	fn := fun

	go fn("func to var goroutine call")

	time.Sleep(time.Second * 3)

	fmt.Println("done..")
}
