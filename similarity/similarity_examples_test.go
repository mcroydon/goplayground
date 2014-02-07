package similarity

import (
	"fmt"
	"github.com/mcroydon/goplayground/similarity"
)

func Example() {
	// Create a similarity engine
	sim := similarity.New()

	// Create some items
	critic1 := "critic1"
	critic2 := "critic2"
	rating1 := similarity.Item{"In a World", 3.5}
	rating2 := similarity.Item{"In a World", 2.0}
	rating3 := similarity.Item{"War Games", 4.5}
	rating4 := similarity.Item{"War Games", 3.0}

	// Add items to the similarity with a string key.
	// In this case we are using the critic's name.
	sim.Add(critic1, rating1)
	sim.Add(critic1, rating3)
	sim.Add(critic2, rating2)
	sim.Add(critic2, rating4)

	// Get similar keys (in this case critics) and the
	// rated items.
	results := sim.SimilarEuclidean("critic1", 5)

	// We can then retrieve the critic's name and similarity score.
	fmt.Println(results[0].Name)
	fmt.Println(results[0].Similarity)
	// Output:
	// critic2
	// 0.32037724101704074
}
