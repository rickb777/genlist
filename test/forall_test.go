package main

import "testing"

func TestForall(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	all1 := things.Forall(func(x Thing) bool {
		return x.Name == "First"
	})

	if all1 {
		t.Errorf("Forall should be false for Name == 'First'")
	}

	all2 := things.Forall(func(x Thing) bool {
		return x.Name == "First" || x.Name == "Second" || x.Name == "Third"
	})

	if !all2 {
		t.Errorf("Forall should be true")
	}

	all3 := ThingSlice{}.Forall(func(x Thing) bool {
		return false
	})

	if !all3 {
		t.Errorf("Forall should evaluate true for empty slices")
	}
}
