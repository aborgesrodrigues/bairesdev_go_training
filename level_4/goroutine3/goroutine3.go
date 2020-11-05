package main

import (
	"fmt"
)

func goroutine(value int, ch chan int) {
	ch <- value
}

func main() {
	ch := make(chan int)

	go goroutine(10, ch)
	go goroutine(100, ch)
	fmt.Println("Channel", <-ch)
	fmt.Println("Channel", <-ch)
}
