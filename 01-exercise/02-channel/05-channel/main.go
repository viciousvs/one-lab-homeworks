package main

import "fmt"

func main() {
	//TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	data := []int{1, 2, 3, 4, 5, 42, -1}

	owner := func(data []int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, v := range data {
				if v < 0 {
					return
				}
				out <- v
			}
		}()
		return out
	}

	ch := owner(data)
	consumer(ch)
}
