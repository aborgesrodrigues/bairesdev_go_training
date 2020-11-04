package main

import (
	"fmt"
)

type queue struct {
	items []int
}

func (q *queue) enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *queue) dequeue() int {
	var removed int = q.items[0]
	q.items = q.items[1:]
	return removed
}

func main() {
	q := queue{items: make([]int, 0)}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q)

	fmt.Println(q.dequeue(), q)
	fmt.Println(q.dequeue(), q)
}
