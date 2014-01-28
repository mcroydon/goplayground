package similarity

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
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
	data  map[string]map[string]Item
}

// Create a new Similarity engine.
func NewSimilarity() *Similarity {
	return &Similarity{data: make(map[string]map[string]Item)}
}

// Add an Item to the engine with a key.
func (sim *Similarity) Add(key string, item Item) {
	sim.mutex.Lock()
	m, ok := sim.data[key]
	if !ok {
		m = make(map[string]Item)
	}
	m[item.Name] = item
	sim.data[key] = m
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

// Dump the data backing a Similarity engine to a Writer.
func (sim *Similarity) WriteJson(w io.Writer) error {
	b, err := json.Marshal(sim.data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// Read the data for a Similarity engine from a Writer and load it
func (sim *Similarity) ReadJson(r io.Reader) error {
	// TODO: Is it idomatic to return underlying errors such as those encountered by ReadFrom or Unmarshal?
	b := new(bytes.Buffer)
	_, err := b.ReadFrom(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b.Bytes(), &sim.data)
	if err != nil {
		return err
	}
	return nil
}

// Returns the Euclidean distance of two keys in our Similarity engine.
func (sim *Similarity) EuclideanDistance(key1 string, key2 string) float64 {
	// Don't compute if either key is missing.
	if sim.data[key1] == nil || sim.data[key2] == nil {
		return -1
	}
	firstItems := sim.data[key1]
	secondItems := sim.data[key2]
	// Find common Items for the two keys
	sum := 0.0
	for _, item := range firstItems {
		secondItem, found := secondItems[item.Name]
		if found {
			sum += math.Pow(item.Value-secondItem.Value, 2)
		}
	}
	return 1 / (1 + math.Sqrt(sum))
}
