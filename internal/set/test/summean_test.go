package main

import "testing"

func TestMean(t *testing.T) {
	others := num1Collection(50, 100, 9, 6, 101, 99)

	if others.Size() != 6 {
		t.Errorf("Size should be 6, got %#v", others.Size())
	}

	sum := others.Sum()

	if sum != 365 {
		t.Errorf("Sum should result in %#v, got %#v", 365, sum)
	}

	mean1 := others.Mean()

	expected := 365.0 / 6.0

	if mean1 != expected {
		t.Errorf("Mean should be %#v, got %#v", expected, mean1)
	}
}
