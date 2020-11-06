package main

import (
	"fmt"
	"sync"
)

func increase() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	var total int

	for r := 0; r < 100; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mutex.Lock()
			total++
			mutex.Unlock()

			fmt.Println(total)
		}()
	}

	wg.Wait()
}

func main() {
	increase()
}
