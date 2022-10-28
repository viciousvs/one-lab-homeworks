package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(50 * time.Millisecond):
					// Simulate work.
					// Report result.
					ch <- data{"123"}
				}
			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
	defer cancel()
	ch := compute(ctx)
	d := <-ch
	fmt.Printf("work complete: %s\n", d)

}
