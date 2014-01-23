// Package tribblestore is a silly package for playing with tribbles, and attempting to contain them.
package tribblestore

import (
	"fmt"
	"math/rand"
)

// Tribble represents a simplistic model of a tribble.
type Tribble struct {
	Color string
}

// NewTribble returns a tribble with a specified color.
func NewTribble(color string) Tribble {
	tribble := Tribble{color}
	return tribble
}

// Eat allows the tribble to eat, which is one of the two things it is best at.
func (t Tribble) Eat() {
	fmt.Printf("The %v tribble goes nom.\n", t.Color)
}

// Reproduce returns a slice containing the offspring of the given tribble.
// Reproducing is the other thing that tribbles are best at.
func (t Tribble) Reproduce() []Tribble {
	count := rand.Intn(10) + 1
	fmt.Printf("The %v tribble just made %v more tribbles.\n", t.Color, count)
	tribbles := make([]Tribble, count, count)
	for i := 0; i < count; i++ {
		tribbles[i] = NewTribble(t.Color)
	}
	return tribbles
}
