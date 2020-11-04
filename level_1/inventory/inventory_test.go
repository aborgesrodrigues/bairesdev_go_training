package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInventory(t *testing.T) {
	i := inventory{products: make(map[int]product)}

	p := product{id: 1, name: "name 1"}
	err := i.add(p)
	assert.Nil(t, err)

	p = product{id: 2, name: "name 2"}
	err = i.add(p)
	assert.Nil(t, err)
}

func TestInventoryErrorsHandling(t *testing.T) {
	i := inventory{products: make(map[int]product)}

	i.add(product{id: 1, name: "name 1"})
	i.add(product{id: 2, name: "name 2"})

	// test not informed id
	err := i.add(product{name: "name 3"})
	assert.Equal(t, err, errors.New("Inform an id to the product"), "they should be equal")

	// test duplicated id
	err = i.add(product{id: 2, name: "name 4"})
	assert.Equal(t, err, errors.New("Duplicated id"), "they should be equal")
}
