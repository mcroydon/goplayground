package similarity

// An item that has been rated.
type Item struct {
	Name  string
	Value float64
}

// Similarity is a similarity storage and retrieval engine.
type Similarity struct {
	data map[string][]Item
}

// Create a new Similarity engine.
func NewSimilarity() *Similarity {
	return &Similarity{make(map[string][]Item)}
}

// Add an Item to the engine with a key.
func (sim *Similarity) Add(key string, item Item) {
	sim.data[key] = append(sim.data[key], item)
}

// Get all the keys in this Similarity.
func (sim *Similarity) Keys() []string {
	keys := make([]string, 0, len(sim.data))
	for k, _ := range(sim.data) {
		keys = append(keys, k)
	}
	return keys
}
