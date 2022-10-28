package main

import (
	"fmt"
)

func main() {
	var countOfData int = 6
	ch := make(chan int, countOfData)

	go func() {
		defer close(ch)

		// TODO: send all iterator values on channel without blocking
		for i := 0; i < countOfData; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
