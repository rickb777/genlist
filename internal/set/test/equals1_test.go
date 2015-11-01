package main

import "testing"

func TestEqualsNum1(t *testing.T) {
	a := num1Collection(50, 100, 9, 7, 100, 99)
	b := num1Collection(50, 100, 9, 100, 7, 99)
	c := num1Collection(50, 100, 9, 7, 42, 99)
	d := num1Collection()

	if !a.Equals(a) {
		t.Errorf("Equals should be true")
	}

	if !a.Equals(b) {
		t.Errorf("Equals should be true")
	}

	if a.Equals(c) {
		t.Errorf("Equals should be false")
	}

	if a.Equals(d) {
		t.Errorf("Equals should be false")
	}

	if c.Equals(a) {
		t.Errorf("Equals should be false")
	}

	if d.Equals(a) {
		t.Errorf("Equals should be false")
	}
}

func TestEqualsThing(t *testing.T) {
	a := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)
	b := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Zzz", 13},
		Thing{"Zzz", 21},
	)
	c := thingCollection()

	if !a.Equals(a) {
		t.Errorf("Equals should be true")
	}

	if a.Equals(b) {
		t.Errorf("Equals should be false")
	}

	if a.Equals(c) {
		t.Errorf("Equals should be false")
	}

	if b.Equals(a) {
		t.Errorf("Equals should be false")
	}

	if c.Equals(a) {
		t.Errorf("Equals should be false")
	}
}
