package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByAgeRange(t *testing.T) {
	tests := []struct {
		ages     []int
		fromAge  int
		toAge    int
		expected []int
	}{
		{ages: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, fromAge: 2, toAge: 8, expected: []int{2, 3, 4, 5, 6, 7, 8}},
		{ages: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, fromAge: 3, toAge: 5, expected: []int{3, 4, 5}},
		{ages: []int{2, 4, 6, 8, 10}, fromAge: 2, toAge: 8, expected: []int{2, 4, 6, 8}},
		{ages: []int{2, 4, 6, 8, 10}, fromAge: 4, toAge: 2, expected: []int{}},
	}

	for _, tt := range tests {
		assert.Equal(t, Equal(filterByAgeRange(tt.fromAge, tt.toAge, tt.ages), tt.expected), true, "they should be equal")
	}
}
