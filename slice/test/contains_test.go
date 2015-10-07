package main

import "testing"

func TestContains(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	has1 := things.Contains(Thing{"Dummy", 9})

	if has1 {
		t.Errorf("Contains should not evaluate true for Name == Dummy")
	}

	has2 := things.Contains(Thing{"Second", -20})

	if !has2 {
		t.Errorf("Contains should evaluate true for Second,-20")
	}

	has3 := ThingSlice{}.Contains(Thing{"A", 1})

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}
