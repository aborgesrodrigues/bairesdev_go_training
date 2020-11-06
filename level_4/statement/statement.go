package main

import (
	"fmt"
	"strings"
	"sync"
)

var allStrings []string = []string{
	"alessandro",
	"borges",
	"string1",
	"aaboss",
}

func findSusbstring(ch chan string, wg *sync.WaitGroup, substring string, partialStrings []string) {
	defer wg.Done()

	for _, sub := range partialStrings {
		if strings.Contains(sub, substring) {
			fmt.Println("encontrou", sub)
			ch <- sub
		}
	}
}

func checkChannel(ch chan string) {
	for ch1 := range ch {
		fmt.Println(ch1)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	var findString = "bo"

	wg.Add(2)
	go findSusbstring(ch, &wg, findString, allStrings[:len(allStrings)/2])
	go findSusbstring(ch, &wg, findString, allStrings[len(allStrings)/2:])

	go checkChannel(ch)
	wg.Wait()
	close(ch)
}
