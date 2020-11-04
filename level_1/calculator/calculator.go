package main

import (
	"errors"
	"fmt"
)

type calculator struct {
	result float32
}

func (c *calculator) add(num1 float32, num2 float32) float32 {
	c.result = num1 + num2
	return c.result
}

func (c *calculator) subtract(num1 float32, num2 float32) float32 {
	c.result = num1 - num2
	return c.result
}

func (c *calculator) multiply(num1 float32, num2 float32) float32 {
	c.result = num1 * num2
	return c.result
}

func (c *calculator) divide(num1 float32, num2 float32) (float32, error) {
	if num2 == 0 {
		//panic("division by zero")
		return 0, errors.New("division by zero")
	}

	c.result = num1 / num2
	return c.result, nil
}

func main() {
	c := calculator{}

	defer fmt.Println("End")

	fmt.Println(c.add(1, 2.2), c.result)
	fmt.Println(c.subtract(40, 25), c.result)
	fmt.Println(c.multiply(3, 6), c.result)

	result1, err := c.divide(10, 2)
	fmt.Println(result1, c.result, err)
	result2, err := c.divide(10, 0)
	fmt.Println(result2, c.result, err)
}
