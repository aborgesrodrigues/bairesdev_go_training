package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgeFilter(t *testing.T) {
	ages := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, age_filter(2, 3, ages), []int{2, 3}, "they should be equal")
	assert.Equal(t, age_filter(5, 8, ages), []int{5, 6, 7, 8}, "they should be equal")
}
