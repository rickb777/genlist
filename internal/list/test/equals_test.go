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

func TestEqualsNum2(t *testing.T) {
	listA := Num2List{ip(50), ip(100), ip(9), ip(7), ip(100), ip(99)}
	listB := Num2List{ip(50), ip(100), ip(9), ip(7), ip(42), ip(99)}
	listC := Num2List{}

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
