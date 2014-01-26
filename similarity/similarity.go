package similarity

import (
	"sync"
)

// An item that has been rated.
type Item struct {
	Name  string
	Value float64
}

// Similarity is a similarity storage and retrieval engine.
type Similarity struct {
	mutex sync.RWMutex
	data  map[string][]Item
}

// Create a new Similarity engine.
func NewSimilarity() *Similarity {
	return &Similarity{data: make(map[string][]Item)}
}

// Add an Item to the engine with a key.
func (sim *Similarity) Add(key string, item Item) {
	sim.mutex.Lock()
	sim.data[key] = append(sim.data[key], item)
	sim.mutex.Unlock()
}

// Get all the keys in this Similarity.
func (sim *Similarity) Keys() []string {
	sim.mutex.RLock()
	keys := make([]string, 0, len(sim.data))
	for k, _ := range sim.data {
		keys = append(keys, k)
	}
	sim.mutex.RUnlock()
	return keys
}
