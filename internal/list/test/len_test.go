package main

import "testing"

func TestLenOther(t *testing.T) {
	things := OtherList{60, -20, 100}

	if things.Len() != 3 {
		t.Errorf("Size should be 3")
	}

	if things.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	if !things.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

	noThings := OtherList{}

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

func TestLenThing(t *testing.T) {
	things := ThingList{
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

	noThings := ThingList{}

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
