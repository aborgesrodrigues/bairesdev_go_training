package main

import "fmt"

type books struct {
	name  string
	price float32
}

type games struct {
	name  string
	price float32
}

type Product interface {
	print()
	applyDescount(float32)
}

func (b *books) print() {
	fmt.Println("Book", b.name, b.price)
}

func (g *games) print() {
	fmt.Println("Game", g.name, g.price)
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

	p.print()
}

func main() {
	b1 := &books{name: "Book 1", price: 10}
	g1 := &games{name: "Game 1", price: 100}

	applyDescount(b1)
	applyDescount(g1)
}
