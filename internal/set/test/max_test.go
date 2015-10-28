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
		Thing{"First", 10},
		Thing{"Second", -10},
		Thing{"Third", 100},
		Thing{"Second", 50},
	)

	m1 := things.Max(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Third", 100}

	if m1 != expected1 {
		t.Errorf("Max Number should return %#v, got %#v", expected1, m1)
	}
}
