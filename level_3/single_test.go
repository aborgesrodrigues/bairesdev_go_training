package main

import "testing"

func TestIntMin(t *testing.T) {
	if got := IntMin(2, 1); got != 1 {
		t.Errorf("IntMin(2, 1) = 1 want 1")
	}
}
