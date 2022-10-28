package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: identify the data race
// fix the issue.

func main() {
	start := time.Now()
	resetCh := make(chan bool)
	t := time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		resetCh <- true
	})
	for time.Since(start) < 5*time.Second {
		<-resetCh
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------

// ==================
// WARNING: DATA RACE
// Read at 0x00c000132018 by goroutine 8:
//   main.main.func1()
//       lab3/01-exercise/05-race/main.go:17 +0xd3

// Previous write at 0x00c000132018 by main goroutine:
//   main.main()
//       lab3/01-exercise/05-race/main.go:15 +0x164

// Goroutine 8 (running) created at:
//   time.goFunc()
//       /usr/lib/go/src/time/sleep.go:176 +0x47
// ==================
// 1.031101653s
// 1.69787255s
// 1.933199843s
// 2.220711436s
// 2.770438186s
// 3.404313178s
// 3.737001578s
// 3.920422622s
// 4.4010764s
// Found 1 data race(s)
// exit status 66
