package main

import (
	"testing"
)

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
