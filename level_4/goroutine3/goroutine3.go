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
	go goroutine(20, ch)
	go goroutine(30, ch)
	go goroutine(40, ch)
	go goroutine(50, ch)

	fmt.Println("Channel", <-ch)
	fmt.Println("Channel", <-ch)
	fmt.Println("Channel", <-ch)
	fmt.Println("Channel", <-ch)
	fmt.Println("Channel", <-ch)
}
