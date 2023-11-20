package main

import "fmt"

func main() {
	valueCh := make(chan int)

	go func(a, b int) {
		valueCh <- a + b
	}(1, 2)

	fmt.Printf("computed value %v\n", <-valueCh)
}
