package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, number := range nums {
			out <- number
		}
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for number := range in {
			out <- number * number
		}
	}()
	return out
}

func main() {
	// set up the pipeline
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	i := 0
	sq := square(generator(numbers...))

	for number := range sq {
		fmt.Printf("%d^2 = %d\n", numbers[i], number)
		i++
	}
	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
