package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) pop() int {
	poped := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]

	return poped
}

func (s *Stack) peek(i int) int {
	for index, item := range s.items {
		if item == i {
			return index
		}
	}
	return -1
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{items: make([]int, 0)}
	fmt.Println(stack.isEmpty())
	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack)
	fmt.Println("Removed ", stack.pop())

	stack.push(4)
	stack.push(5)
	fmt.Println(stack)

	fmt.Println("Item 4", "Index", stack.peek(4))
	fmt.Println(stack.isEmpty())
}
