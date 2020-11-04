package main

type User struct {
	name string
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
