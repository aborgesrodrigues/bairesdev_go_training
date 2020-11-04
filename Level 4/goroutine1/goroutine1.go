package main

import "fmt"

func goroutine() {
	fmt.Println("Hello World")
}

func main() {
	go goroutine()
	fmt.Println("main function")
}
