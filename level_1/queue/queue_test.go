package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := queue{items: make([]int, 0)}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	assert.Equal(t, q.items, []int{1, 2, 3}, "they should be equal")
}

func TestDequeue(t *testing.T) {
	q := queue{items: []int{1, 2, 3, 4}}

	assert.Equal(t, q.dequeue(), 1, "they should be equal")
	assert.Equal(t, q.items, []int{2, 3, 4}, "they should be equal")

	assert.Equal(t, q.dequeue(), 2, "they should be equal")
	assert.Equal(t, q.items, []int{3, 4}, "they should be equal")
}
