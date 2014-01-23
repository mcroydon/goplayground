package main

import (
	"fmt"
	"github.com/mcroydon/goplayground/tribblestore"
	"math/rand"
	"time"
)

func count(size int) {
	fmt.Printf("\nThere are now %v tribbles in the cargo hold.\n", size)
}

func main() {
	// In a world...
	rand.Seed(time.Now().UTC().UnixNano())

	// Create our first tribble.
	tribble := tribblestore.NewTribble("gray")
	fmt.Println("Is that a tribble?\n")

	// Give it some food.
	tribble.Eat()

	// This is inevitable.
	tribbles := tribble.Reproduce()

	// Time goes by.
	for i := 0; i < 100; i++ {
		for _, t := range tribbles {
			t.Eat()
		}
		// Note that this is not an accurate simulation as more than one
		// tribble is likely reproducing at any given moment.
		babytribbles := tribbles[rand.Intn(len(tribbles))].Reproduce()
		tribbles = append(tribbles, babytribbles...)
		count(len(tribbles))
	}


}
