package main

import (
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMin(t *testing.T) {
	if got := IntMin(2, 1); got != 1 {
		t.Errorf("IntMin(2, 1) = 1 want 1")
	}
}
