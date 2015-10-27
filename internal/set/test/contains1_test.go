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
		Thing{"First", 60},
		Thing{"Second", -20},
		Thing{"Third", 100},
	)

	has1 := things.Contains(Thing{"Dummy", 9})

	if has1 {
		t.Errorf("Contains should not evaluate true for Name == Dummy")
	}

	has2 := things.Contains(Thing{"Second", -20})

	if !has2 {
		t.Errorf("Contains should evaluate true for Second,-20")
	}

	has3 := thingCollection().Contains(Thing{"A", 1})

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}
