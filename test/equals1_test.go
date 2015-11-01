package main

import "testing"

func TestEqualsFoo1(t *testing.T) {
	listA := NewFoo1List(50, 100, 9, 7, 100, 99)
	setA := NewFoo1Set(50, 100, 9, 7, 100, 99)
	setB := NewFoo1Set(50, 100, 9, 100, 7, 99)

	if !listA.Equals(listA) {
		t.Errorf("Equals should be true")
	}

	if listA.Equals(setA) {
		t.Errorf("Equals should be false")
	}

	if setA.Equals(listA) {
		t.Errorf("Equals should be false")
	}

	if !setA.Equals(setA) {
		t.Errorf("Equals should be true")
	}

	if !setB.Equals(setA) {
		t.Errorf("Equals should be true")
	}
}
