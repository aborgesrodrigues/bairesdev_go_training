package main

import (
	"fmt"
	"sync"
	"time"
)

func increase() {
	var mutex = &sync.Mutex{}
	total := 0
	for r := 0; r < 100; r++ {
		go func() {
			mutex.Lock()
			total++
			fmt.Println(total)
			mutex.Unlock()
			// Wait a bit between reads.
			time.Sleep(time.Millisecond)
		}()
	}
}

func main() {
	// Doubt about the example https://tour.golang.org/methods/14
	increase()
}
