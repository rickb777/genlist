package main

import "testing"

func TestMean(t *testing.T) {
	others := num1Collection(50, 100, 9, 7, 100, 99)

	mean1 := others.Mean()

	avg1 := 365.0 / 6.0

	if mean1 != avg1 {
		t.Errorf("Mean should be %#v, got %#v", avg1, mean1)
	}
}
