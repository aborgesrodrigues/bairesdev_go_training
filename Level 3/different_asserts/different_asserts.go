package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	name string
}

func TestIntMin(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var u User
	assert.Nil(t, u)

	// assert for not nil (good when you expect something)
	var u2 = User{name: "Alessandro"}
	if assert.NotNil(t, u2) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Name", u2.name)

	}
}
