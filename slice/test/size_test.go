package main

import "testing"

func TestSize(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	if things.Len() != 3 {
		t.Errorf("Size should be 3")
	}

	if things.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	if !things.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

	noThings := ThingSlice{}

	if noThings.Len() != 0 {
		t.Errorf("Size should be 0")
	}

	if !noThings.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}

	if noThings.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}
}
