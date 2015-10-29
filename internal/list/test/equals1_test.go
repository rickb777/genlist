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
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)
	listB := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Zzz", 13},
		Thing{"Zzz", 21},
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
