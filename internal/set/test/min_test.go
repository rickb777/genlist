package main

import "testing"

func TestMinOrdered(t *testing.T) {
	stuff := num1Collection(50, 100, 9, 7, 100, 99)

	min1 := stuff.Min()
	m1 := Num1(7)

	if min1 != m1 {
		t.Errorf("Min should return %#v, got %#v", m1, min1)
	}
}

func TestMinNotOrdered(t *testing.T) {
	things := thingCollection(
		Thing{"First", 10},
		Thing{"Second", -10},
		Thing{"Third", 100},
		Thing{"Second", 50},
	)

	m1 := things.Min(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Second", -10}

	if m1 != expected1 {
		t.Errorf("Min Number should return %#v, got %#v", expected1, m1)
	}
}
