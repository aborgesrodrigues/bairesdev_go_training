package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	w float32
	h float32
}

type Circle struct {
	r float32
}

type Shape interface {
	Area() float32
	Perimeter() float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.r * c.r
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) Area() float32 {
	return r.w * r.h
}

func (r *Rectangle) Perimeter() float32 {
	return (2 * r.w) + (2 * r.h)
}

func run(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func main() {
	var shapes = make([]Shape, 0)
	fmt.Println(shapes)

	c1 := &Circle{r: 3}
	shapes = append(shapes, c1)

	r1 := &Rectangle{w: 3, h: 4}
	shapes = append(shapes, r1)

	for _, shape := range shapes {
		run(shape)
	}
}
