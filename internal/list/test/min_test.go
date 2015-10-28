package main

import "testing"

func TestMinOrdered(t *testing.T) {
	others := num1Collection(50, 100, 9, 7, 100, 99)

	min1 := others.Min()
	m1 := Num1(7)

	if min1 != m1 {
		t.Errorf("Min should return %#v, got %#v", m1, min1)
	}
}

func TestMinNotOrdered(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
		{"Fourth", 10},
	}

	max1 := things.Min(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Second", -20}

	if max1 != expected1 {
		t.Errorf("Min Number should return %#v, got %#v", expected1, max1)
	}
}
