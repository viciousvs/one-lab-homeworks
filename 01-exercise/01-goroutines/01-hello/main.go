package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func asyncFunc(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.
	// goroutine function call
	wg.Add(1)
	go asyncFunc("function call")

	// goroutine with anonymous function
	wg.Add(1)
	go func() {
		defer wg.Done()
		asyncFunc("anonymous function call")
	}()

	// goroutine with function value call
	wg.Add(1)
	go func(fn func(string), msg string) {
		defer wg.Done()
		fn(msg)
	}(asyncFunc, "funcion value call")

	// wait for goroutines to end
	wg.Wait()
	fmt.Println("done..")
}
