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
		Thing{"Fee", 3},
		Thing{"Fie", 2},
		Thing{"Foe", 1},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	max1 := things.Min(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Foe", 1}

	if max1 != expected1 {
		t.Errorf("Min Number should return %#v, got %#v", expected1, max1)
	}
}
