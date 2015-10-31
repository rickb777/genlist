package main

import "testing"

func TestMinByNum(t *testing.T) {
	things := num1Collection(50, 100, -20, 7, 100, 99)

	min1 := things.MinBy(func(a, b Num1) bool {
		return a < b
	})

	if min1 != -20 {
		t.Errorf("MinBy should return %#v, got %#v", -20, min1)
	}
}

func TestMinByThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 3},
		Thing{"Fie", 2},
		Thing{"Foe", 1},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	min1 := things.MinBy(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	expected1 := Thing{"Foe", 1}
	if min1 != expected1 {
		t.Errorf("MinBy should return %#v, got %#v", expected1, min1)
	}
}
