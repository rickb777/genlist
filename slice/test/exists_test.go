package main

import "testing"

func TestExists(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	any1 := things.Exists(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if any1 {
		t.Errorf("Exists should not evaluate true for Name == Dummy")
	}

	any2 := things.Exists(func(x Thing) bool {
		return x.Number > 50
	})

	if !any2 {
		t.Errorf("Exists should evaluate true for Number > 50")
	}

	any3 := ThingSlice{}.Exists(func(x Thing) bool {
		return true
	})

	if any3 {
		t.Errorf("Exists should evaluate false for empty slices")
	}
}
