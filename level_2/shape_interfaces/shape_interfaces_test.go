package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapes(t *testing.T) {
	var shapes = make([]Shape, 0)

	c1 := &Circle{r: 3}
	shapes = append(shapes, c1)

	r1 := &Rectangle{w: 3, h: 4}
	shapes = append(shapes, r1)

	for _, shape := range shapes {
		switch shape.(type) {
		case *Circle:
			assert.Equal(t, shape.Area(), float32(28.274334), "they should be equal")
			assert.Equal(t, shape.Perimeter(), float32(18.849556), "they should be equal")
		case *Rectangle:
			assert.Equal(t, shape.Area(), float32(12), "they should be equal")
			assert.Equal(t, shape.Perimeter(), float32(14), "they should be equal")
		}
	}
}
