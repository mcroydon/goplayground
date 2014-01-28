package similarity

import (
	"bytes"
	"encoding/json"
	"io"
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
