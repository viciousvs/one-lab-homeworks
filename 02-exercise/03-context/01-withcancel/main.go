package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumber() int {
	return rand.Intn(10)
}

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					out <- GenerateNumber()
				}
			}
		}()
		return out
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())
	i := 0

	for number := range generator(ctx) {
		fmt.Println(number)
		if i == 50 {
			cancel()
		}
		i++
	}
}
