package counting

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var WorkerCount = runtime.GOMAXPROCS(runtime.NumCPU())

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	fmt.Println(WorkerCount)
	var sum int64
	var wg sync.WaitGroup

	result := make(chan int64)
	defer close(result)

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for s := range result {
			sum += s
			i++
			if i == WorkerCount {
				return
			}
		}
	}()

	workersWorkCount := len(numbers) / WorkerCount

	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go func(ind int) {
			defer wg.Done()
			log.Printf("worker by index:%d done", ind)
			offset := ind * workersWorkCount
			worker(result, numbers[offset:offset+workersWorkCount])
		}(i)
	}
	wg.Wait()
	return sum
}

func worker(result chan int64, data []int) {
	sum := Add(data)
	result <- sum
}
