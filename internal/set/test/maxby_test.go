package main

import "testing"

func TestMaxByNum(t *testing.T) {
	things := num1Collection(50, 100, -20, 7, 100, 99)

	max := things.MaxBy(func(a, b Num1) bool {
		return a < b
	})

	if max != 100 {
		t.Errorf("MaxBy Number should return %#v, got %#v", 100, max)
	}
}

func TestMaxByThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 21},
		Thing{"Bam", 13},
		Thing{"Bam", 8},
	)

	max := things.MaxBy(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Boo", 21}
	if max != expected1 {
		t.Errorf("MaxBy Number should return %#v, got %#v", expected1, max)
	}
}
