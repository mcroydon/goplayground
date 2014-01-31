package similarity

import (
	"bytes"
	"testing"
)

var (
	jsondata = []byte(`{"critic1":{"In a World":{"Name":"In a World","Value":3.5},"War Games":{"Name":"War Games","Value":4.5}},
		"critic2":{"In a World":{"Name":"In a World","Value":2},"War Games":{"Name":"War Games","Value":3}}}`)
)

func populateSim() *Similarity {
	critic1 := "critic1"
	critic2 := "critic2"
	rating1 := Item{"In a World", 3.5}
	rating2 := Item{"In a World", 2.0}
	rating3 := Item{"War Games", 4.5}
	rating4 := Item{"War Games", 3.0}

	sim := NewSimilarity()
	sim.Add(critic1, rating1)
	sim.Add(critic2, rating2)
	sim.Add(critic1, rating3)
	sim.Add(critic2, rating4)
	return sim
}

func TestCreate(t *testing.T) {

	sim := populateSim()
	keys := sim.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected %v keys, had %v.", 2, len(keys))
	}

	buf := new(bytes.Buffer)

	sim.WriteJson(buf)
}

func TestReadJson(t *testing.T) {
	sim := NewSimilarity()
	err := sim.ReadJson(bytes.NewBuffer(jsondata))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	keys := sim.Keys()
	expectedKeys := []string{"critic1", "critic2"}
	if len(keys) == 2 && keys[0] != "critic1" || keys[1] != "critic2" {
		t.Errorf("Expected keys %v found keys %v.", expectedKeys, keys)
	}
}

func TestWriteJson(t *testing.T) {
	sim := populateSim()
	b := new(bytes.Buffer)
	err := sim.WriteJson(b)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !bytes.Contains(b.Bytes(), []byte("In a World")) {
		t.Errorf("Output does not contain expected value: %v", b.Bytes())
	}
}

func TestGet(t *testing.T) {
	sim := populateSim()
	key := "critic1"
	name := "In a World"
	m := sim.Get(key)
	value := 3.5
	if m[name].Value != value {
		t.Errorf("Expected a rating of %v for %v %v.", value, name, key)
	}
}

func TestSimilarity(t *testing.T) {
	sim := populateSim()
	score := sim.EuclideanDistance("critic1", "critic2")
	expected := 0.32037724101704074
	if score != expected {
		t.Errorf("Found unexpected similarity %v, (expected %v).", score, expected)
	}
	results := sim.SimilarEuclidean("critic1", 5)
	if len(results) != 1 {
		t.Errorf("Expected 1 result found %v.", len(results))
	}
	if results[0].Name != "critic2" && results[0].Similarity != expected {
		t.Errorf("Found unexpected score %v for %v.", results[0].Similarity, results[0].Name)
	}
}
