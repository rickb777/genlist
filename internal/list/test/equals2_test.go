package main

import "testing"

func TestEqualsNum2(t *testing.T) {
	listA := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(100), ip(99))
	listB := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(42), ip(99))
	listC := num2Collection()

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
