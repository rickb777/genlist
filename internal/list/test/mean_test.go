package main

import "testing"

func TestMean(t *testing.T) {
	others := Num1List{50, 100, 9, 7, 100, 99}

	mean1, err := others.Mean()

	if err != nil {
		t.Errorf("Mean should succeed")
	}

	avg1 := Num1(60)

	if mean1 != avg1 {
		t.Errorf("Mean should be %v, got %v", avg1, mean1)
	}

	mean2, err := Num1List{}.Mean()

	if err == nil || mean2 != 0 {
		t.Errorf("Mean should fail on empty list")
	}
}
