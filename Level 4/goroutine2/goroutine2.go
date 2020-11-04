package main

import (
	"fmt"
	"time"
)

func goroutine() {
	fmt.Println("Hello World")
}

func main() {
	go goroutine()
	fmt.Println("main function")
	time.Sleep(2 * time.Second)
}
