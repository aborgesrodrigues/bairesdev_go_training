package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntMin(t *testing.T) {
	if got := IntMin(2, 1); got != 1 {
		t.Errorf("IntMin(2, 1) = 1 want 1")
	}
}

func TestMultipleIntMin(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{a: 1, b: 2, result: 1},
		{a: 2, b: 1, result: 1},
		{a: 3, b: 7, result: 3},
		{a: 4, b: 2, result: 2},
	}
	for _, tt := range tests {
		if got := IntMin(tt.a, tt.b); got != tt.result {
			t.Errorf("IntMin(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.result)
		}
	}
}

func TestDifferentAssertsIntMin(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, nil)

	// assert for not nil (good when you expect something)
	var u2 = User{name: "Alessandro"}
	if assert.NotNil(t, u2) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Alessandro", u2.name)

	}
}
