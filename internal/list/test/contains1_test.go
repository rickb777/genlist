package main

import (
	"testing"
)

func TestContainsNum1(t *testing.T) {
	things := num1Collection(50, 100, 9, 7, 100, 99)

	has1 := things.Contains(3)

	if has1 {
		t.Errorf("Contains should not evaluate true for 3")
	}

	has2 := things.Contains(7)

	if !has2 {
		t.Errorf("Contains should evaluate true for 7")
	}

	has3 := num1Collection().Contains(1)

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}

func TestContainsThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	has1 := things.Contains(Thing{"Dummy", 9})

	if has1 {
		t.Errorf("Contains should not evaluate true for Name == Dummy")
	}

	has2 := things.Contains(Thing{"Fie", 2})

	if !has2 {
		t.Errorf("Contains should evaluate true for Fie,2")
	}

	has3 := thingCollection().Contains(Thing{"A", 1})

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}
