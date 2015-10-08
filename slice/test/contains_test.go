package main

import "testing"

func TestContainsOther(t *testing.T) {
	things := OtherSlice{50, 100, 9, 7, 100, 99}

	has1 := things.Contains(3)

	if has1 {
		t.Errorf("Contains should not evaluate true for 3")
	}

	has2 := things.Contains(7)

	if !has2 {
		t.Errorf("Contains should evaluate true for 7")
	}

	has3 := OtherSlice{}.Contains(1)

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}

func TestContainsThing(t *testing.T) {
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
