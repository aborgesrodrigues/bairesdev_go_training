package main

import (
	"errors"
	"fmt"
)

type product struct {
	id   int
	name string
}

type inventory struct {
	products map[int]product
}

/*
	Add an product to the inventory
*/
func (i *inventory) add(p product) error {
	// Doubt error handling
	if p.id == 0 {
		//panic("Inform an id to the product")
		return errors.New("Inform an id to the product")
	} else if _, ok := i.products[p.id]; ok {
		return errors.New("Duplicated id")
		//panic("Duplicated id")
	}
	i.products[p.id] = p
	return nil
}

/*
	Print the object and the error if happened
*/
func print(p product, err error) {
	if err != nil {
		fmt.Println("Error", err, p.name)
	} else {
		fmt.Println("Success", p.id, p.name)
	}
}

func main() {
	i := inventory{products: make(map[int]product)}
	fmt.Println(i)

	p := product{id: 1, name: "name 1"}
	err := i.add(p)
	print(p, err)

	p = product{id: 2, name: "name 2"}
	err = i.add(p)
	print(p, err)

	p = product{name: "name 3"}
	err = i.add(p)
	print(p, err)

	p = product{id: 2, name: "name 4"}
	err = i.add(p)
	print(p, err)
}
