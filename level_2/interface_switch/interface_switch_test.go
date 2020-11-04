package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchInterface(t *testing.T) {
	u := &User{"Matt", "Aimonetti"}
	c := &Customer{42, "Francesc"}

	assert.Equal(t, GetName(u), "User Matt Aimonetti", "they should be equal")
	assert.Equal(t, GetName(c), "Customer Francesc", "they should be equal")
}
