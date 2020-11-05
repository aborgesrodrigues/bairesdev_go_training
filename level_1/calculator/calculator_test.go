package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorAdd(t *testing.T) {
	c := calculator{}

	assert.Equal(t, c.add(1, 2.2), float32(3.2), "they should be equal")
}

func TestCalculatorSubtract(t *testing.T) {
	c := calculator{}

	assert.Equal(t, c.subtract(40, 25), float32(15), "they should be equal")
}

func TestCalculatorMultiply(t *testing.T) {
	c := calculator{}

	assert.Equal(t, c.multiply(3, 6), float32(18), "they should be equal")
}

func TestCalculatorDivide(t *testing.T) {
	c := calculator{}

	result1, _ := c.divide(10, 2)
	assert.Equal(t, result1, float32(5), "they should be equal")

	// test the division by zero error handling
	_, err := c.divide(10, 0)
	assert.Equal(t, err, errors.New("division by zero"), "they should be equal")
}
