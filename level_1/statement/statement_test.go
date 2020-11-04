package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stack := Stack{items: make([]int, 0)}

	stack.push(1)
	stack.push(2)
	stack.push(3)

	assert.Equal(t, stack.items, []int{1, 2, 3}, "they should be equal")
}

func TestPop(t *testing.T) {
	stack := Stack{items: []int{1, 2, 3}}

	assert.Equal(t, stack.pop(), 3, "they should be equal")
	assert.Equal(t, stack.items, []int{1, 2}, "they should be equal")
}

func TestPeek(t *testing.T) {
	stack := Stack{items: []int{1, 2, 3}}

	assert.Equal(t, stack.peek(2), 1, "they should be equal")
}

func TestIsEmpty(t *testing.T) {
	stack := Stack{items: make([]int, 0)}
	assert.Equal(t, stack.isEmpty(), true, "they should be equal")

	stack.push(1)
	assert.Equal(t, stack.isEmpty(), false, "they should be equal")
}
