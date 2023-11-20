package main

import "fmt"

// DONE: Implement relaying of message with Channel Direction

func genMsg(ch1 chan interface{}) {
	const generatedMsg = "Generated message from genMsg"
	// send message on ch1
	ch1 <- generatedMsg
}

func relayMsg(ch1, ch2 chan interface{}) {
	defer close(ch1)
	defer close(ch2)

	// recv message on ch1
	msg := <-ch1
	// send it on ch2
	ch2 <- msg
}

func main() {
	// create ch1 and ch2
	var (
		ch1 = make(chan interface{})
		ch2 = make(chan interface{})
	)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	// recv message on ch2
	fmt.Println(<-ch2)
}
