package main

import (
	"fmt"
)

type Trumpeter struct {
	name string
}

type Violinist struct {
	name string
}

type MusicalPlayer interface {
	Greetings() string
}

func (c *Trumpeter) Greetings() string {
	return "Greetings Trumpeter"
}

func (c *Violinist) Greetings() string {
	return "Greetings Violinist"
}

func main() {
	var musicians = make([]MusicalPlayer, 0)
	fmt.Println(musicians)

	v1 := &Violinist{name: "v1"}
	musicians = append(musicians, v1)

	t1 := &Trumpeter{name: "t1"}
	musicians = append(musicians, t1)

	for i, musician := range musicians {
		fmt.Println(i, musician.Greetings())
	}
}
