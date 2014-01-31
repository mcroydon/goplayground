package similarity

import (
	"math/rand"
	"fmt"
	"testing"
	"time"
)

func BenchmarkLoad(b *testing.B) {
	sim := NewSimilarity()
	for i := 0; i < b.N; i++ {
		name := fmt.Sprintf("critic%v", i)
		item := Item{name, rand.NormFloat64() * 5 + 3.5}
		sim.Add(name, item)
	}
}

func BenchmarkSimilarity(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())
	sim := NewSimilarity()
	// Add a first user
	for k := 0; k <= 10; k++ {
		sim.Add("critic0", Item{fmt.Sprintf("item%v", k), 2.5})
	}

	// Add additional users
	for i := 0; i < b.N; i++ {
		for j := 0; j <= 10; j++ {
			name := fmt.Sprintf("item%v", j)
			item := Item{name, rand.Float64() * 5 + 1}
			sim.Add(fmt.Sprintf("critic%v", i + 1), item)
		}
	}
	b.ResetTimer()
	results := sim.SimilarEuclidean("critic0", 10)
	if len(results) < 1 {
		b.Errorf("Did not get any results.")
	} else {
		b.Logf("First result is %v with a score of %v.", results[0].Name, results[0].Similarity)
		for _, item := range sim.Get(results[0].Name) {
			b.Logf("%v: %v", item.Name, item.Value)
		}
	}
}
