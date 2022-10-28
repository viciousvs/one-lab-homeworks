package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// TODO: multiplex recv on channel - ch1, ch2
	select {
	case s := <-ch1:
		recvFrom(s)
	case s := <-ch2:
		recvFrom(s)
	}
}

func recvFrom(ch string) {
	fmt.Println("recv from " + ch)
}
