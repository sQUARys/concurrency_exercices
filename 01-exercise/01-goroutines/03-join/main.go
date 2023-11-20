package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var (
		data int64
		wg   sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.AddInt64(&data, 1)
	}()

	wg.Wait()

	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
