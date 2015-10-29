package main

import "testing"

func TestMaxOrdered(t *testing.T) {
	others := num1Collection(50, 100, 9, 7, 100, 99)

	max1 := others.Max()
	m1 := Num1(100)

	if max1 != m1 {
		t.Errorf("Max should return %#v, got %#v", m1, max1)
	}
}

func TestMaxNotOrdered(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 21},
		Thing{"Bam", 13},
		Thing{"Bam", 8},
	)

	max1 := things.Max(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Boo", 21}

	if max1 != expected1 {
		t.Errorf("Max Number should return %#v, got %#v", expected1, max1)
	}
}
