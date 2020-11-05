package main

import (
	"fmt"
)

type books struct {
	name  string
	price float32
}

type games struct {
	name  string
	price float32
}

type Product interface {
	String() string
	applyDescount(float32)
}

func (b *books) String() string {
	return fmt.Sprintf("Book %s %f", b.name, b.price)
	//"Book " + b.name + " " + strconv.FormatFloat(float64(b.price), 'E', -1, 64)
}

func (g *games) String() string {
	return fmt.Sprintf("Game %s %f", g.name, g.price)
	//"Game " + g.name + " " + strconv.FormatFloat(float64(g.price), 'E', -1, 64)
}

func (b *books) applyDescount(percentDescount float32) {
	b.price = b.price * (1 - percentDescount/100)
}

func (g *games) applyDescount(percentDescount float32) {
	g.price = g.price * (1 - percentDescount/100)
}

func applyDescount(p Product) {
	switch p.(type) {
	case *books:
		p.applyDescount(10)
	case *games:
		p.applyDescount(20)
	}

	fmt.Println(p)
}

func main() {
	b1 := &books{name: "Book 1", price: 10}
	g1 := &games{name: "Game 1", price: 100}

	applyDescount(b1)
	applyDescount(g1)
}
