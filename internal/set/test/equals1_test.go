package main

import "testing"

func TestEqualsNum1(t *testing.T) {
	listA := num1Collection(50, 100, 9, 7, 100, 99)
	listB := num1Collection(50, 100, 9, 7, 42, 99)
	listC := num1Collection()

	if !listA.Equals(listA) {
		t.Errorf("Equals should be true")
	}

	if listA.Equals(listB) {
		t.Errorf("Equals should be false")
	}

	if listA.Equals(listC) {
		t.Errorf("Equals should be false")
	}

	if listB.Equals(listA) {
		t.Errorf("Equals should be false")
	}

	if listC.Equals(listA) {
		t.Errorf("Equals should be false")
	}
}

func TestEqualsThing(t *testing.T) {
	listA := thingCollection(
		Thing{"First", 60},
		Thing{"Second", -20},
		Thing{"Third", 100},
	)
	listB := thingCollection(
		Thing{"First", 60},
		Thing{"Second", -20},
		Thing{"Fourth", 10},
	)
	listC := thingCollection()

	if !listA.Equals(listA) {
		t.Errorf("Equals should be true")
	}

	if listA.Equals(listB) {
		t.Errorf("Equals should be false")
	}

	if listA.Equals(listC) {
		t.Errorf("Equals should be false")
	}

	if listB.Equals(listA) {
		t.Errorf("Equals should be false")
	}

	if listC.Equals(listA) {
		t.Errorf("Equals should be false")
	}
}
