package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	data int
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
