package main

import (
	"fmt"
)

// TODO: Implement relaying of message with Channel Direction

func genMsg(msgs []string) <-chan string {
	// send message on ch1
	send := make(chan string)
	go func() {
		defer close(send)
		for _, msg := range msgs {
			send <- msg
		}
	}()
	return send
}

func relayMsg(in <-chan string) <-chan string {
	// recv message on ch1
	// send it on ch2
	out := make(chan string)
	go func() {
		defer close(out)
		for msg := range in {
			out <- "relay " + msg
		}
	}()
	return out
}

func main() {
	// create ch1 and ch2
	// spine goroutine genMsg and relayMsg
	msgs := []string{"hello", "world", "Salem", "alem", "42"}
	genCh := genMsg(msgs)
	relayCh := relayMsg(genCh)
	// recv message on ch2
	for v := range relayCh {
		fmt.Println(v)
	}
}
