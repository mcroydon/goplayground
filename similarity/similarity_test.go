package similarity

import "testing"

func TestLoad(t *testing.T) {
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

	keys := sim.Keys()
	t.Log(keys)
	if len(keys) != 2 {
		t.Errorf("Expected %v keys, had %v.", 2, len(keys))
	}
}
