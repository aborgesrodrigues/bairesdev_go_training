package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSusbstring(t *testing.T) {
	tests := []struct {
		search   string
		expected string
	}{
		{search: "ss", expected: "alessandro"},
		{search: "bo", expected: "borgess"},
	}

	ch := make(chan string)

	for _, tt := range tests {
		go findSusbstring(ch, nil, tt.search, allStrings[:len(allStrings)/2])

		assert.Equal(t, <-ch, tt.expected, "they should be equal")
	}
}
