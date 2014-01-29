package similarity

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
	"sync"
	"sort"
)

// An item that has been rated.
type Item struct {
	Name  string
	Value float64
}

// The result of a similarity search. Name can be used to retrieve full information
// from the store and Similarity is the similarity score.
type Result struct {
	Name       string
	Similarity float64
}

// A function used to compare the similarity of two keys in the Similarity engine.
type comparison func(key1 string, key2 string) (distance float64)

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
	defer sim.mutex.Unlock()
	m, ok := sim.data[key]
	if !ok {
		m = make(map[string]Item)
	}
	m[item.Name] = item
	sim.data[key] = m
}

// Get all the keys in this Similarity.
func (sim *Similarity) Keys() []string {
	sim.mutex.RLock()
	defer sim.mutex.RUnlock()
	keys := make([]string, 0, len(sim.data))
	for k, _ := range sim.data {
		keys = append(keys, k)
	}
	return keys
}

// Dump the data backing a Similarity engine to a Writer.
func (sim *Similarity) WriteJson(w io.Writer) error {
	sim.mutex.RLock()
	defer sim.mutex.RUnlock()
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
	sim.mutex.Lock()
	defer sim.mutex.Unlock()
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
	sim.mutex.RLock()
	defer sim.mutex.RUnlock()
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

// Find similar keys using Euclidean distance comparison.
func (sim *Similarity) SimilarEuclidean(key string, limit int) []Result {
	return sim.Similar(key, limit, sim.EuclideanDistance)
}

// Find similar keys using the provided distance comparison.
func (sim *Similarity) Similar(key string, limit int, distance comparison) []Result {
	results := make([]Result, 0)
	sim.mutex.RLock()
	defer sim.mutex.RUnlock()
	for k, _ := range sim.data {
		if k == key {
			// Don't check ourselves.
			continue
		}
		score := distance(key, k)
		// TODO: replace -1 return value with an error in EuclideanDistance.
		if score != -1 {
			results = append(results, Result{k, score})
		}
	}
	sort.Sort(bySimilarity(results))
	if len(results) > limit {
		return results[:limit]
	} else {
		return results
	}

}

// byScore implements sort.Interface for []Result allowing us to sort results by score,
// sorting higher scoring results ahead of lower scoring results.
type bySimilarity []Result

func (a bySimilarity) Len() int           { return len(a) }
func (a bySimilarity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySimilarity) Less(i, j int) bool { return a[i].Similarity > a[j].Similarity }
