package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyDescount(t *testing.T) {
	b1 := &books{name: "Book 1", price: 10}
	g1 := &games{name: "Game 1", price: 100}

	applyDescount(b1)
	assert.Equal(t, b1.price, float32(9), "they should be equal")

	applyDescount(g1)
	assert.Equal(t, g1.price, float32(80), "they should be equal")
}

func TestString(t *testing.T) {
	b1 := &books{name: "Book 1", price: 10}
	g1 := &games{name: "Game 1", price: 100}

	assert.Equal(t, b1.String(), "Book Book 1 10.000000", "they should be equal")

	assert.Equal(t, g1.String(), "Game Game 1 100.000000", "they should be equal")
}
