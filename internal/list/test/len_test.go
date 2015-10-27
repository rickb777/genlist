package main

import "testing"

func TestLenNum(t *testing.T) {
	things := num1Collection(60, -20, 100)

	if things.Size() != 3 {
		t.Errorf("Size should be 3")
	}

	if things.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	if !things.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

	if things.ToList().Size() != 3 {
		t.Errorf("Size should be 3")
	}

	noThings := num1Collection()

	if noThings.Size() != 0 {
		t.Errorf("Size should be 0")
	}

	if !noThings.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}

	if noThings.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}

	if noThings.ToList().Size() != 0 {
		t.Errorf("Size should be 0")
	}
}

func TestLenThing(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	if things.Size() != 3 {
		t.Errorf("Size should be 3")
	}

	if things.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	if !things.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

	if things.ToList().Size() != 3 {
		t.Errorf("Size should be 3")
	}

	noThings := ThingList{}

	if noThings.Size() != 0 {
		t.Errorf("Size should be 0")
	}

	if !noThings.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}

	if noThings.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}

	if noThings.ToList().Size() != 0 {
		t.Errorf("Size should be 0")
	}
}
