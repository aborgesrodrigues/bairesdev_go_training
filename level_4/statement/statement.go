package main

import (
	"fmt"
	"strings"
	"sync"
)

var allStrings []string = []string{
	"alessandro",
	"borgess",
	"string1",
	"aaboss",
	"string2",
	"aaboss222",
}

func findSusbstring(ch chan string, wg *sync.WaitGroup, substring string, partialStrings []string) {
	defer wg.Done()

	for _, sub := range partialStrings {
		if strings.Contains(sub, substring) {
			ch <- sub
		}
	}
}

func checkChannel(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch1 := range ch {
		fmt.Println(ch1)
	}
}

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	ch := make(chan string, 6)
	var findString = "ss"

	wg.Add(3)

	go findSusbstring(ch, &wg, findString, allStrings[:2])
	go findSusbstring(ch, &wg, findString, allStrings[2:4])
	go findSusbstring(ch, &wg, findString, allStrings[4:])

	wg2.Add(1)
	go checkChannel(ch, &wg2)
	wg.Wait()
	close(ch)
	wg2.Wait()
}
