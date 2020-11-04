package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMusicalPlayerGreetings(t *testing.T) {
	var musicians = make([]MusicalPlayer, 0)

	v1 := &Violinist{name: "v1"}
	musicians = append(musicians, v1)

	t1 := &Trumpeter{name: "t1"}
	musicians = append(musicians, t1)

	for _, musician := range musicians {
		switch musician.(type) {
		case *Trumpeter:
			assert.Equal(t, musician.Greetings(), "Greetings Trumpeter", "they should be equal")
		case *Violinist:
			assert.Equal(t, musician.Greetings(), "Greetings Violinist", "they should be equal")
		}
	}
}
