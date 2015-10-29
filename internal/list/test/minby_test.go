package main

import "testing"

func TestMinByNum(t *testing.T) {
	things := num1Collection(50, 100, -20, 7, 100, 99)

	min1, err1 := things.MinBy(func(a, b Num1) bool {
		return a < b
	})

	if err1 != nil {
		t.Errorf("MinBy Number should succeed")
	}

	if min1 != -20 {
		t.Errorf("MinBy Number should return %#v, got %#v", -20, min1)
	}

	_, err2 := num1Collection().MinBy(func(a, b Num1) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MinBy Number should fail on empty list")
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

	min1, err1 := things.MinBy(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	if err1 != nil {
		t.Errorf("MinBy Number should succeed")
	}

	expected1 := Thing{"Foe", 1}
	if min1 != expected1 {
		t.Errorf("MinBy Number should return %#v, got %#v", expected1, min1)
	}

	_, err2 := ThingList{}.MinBy(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MinBy Number should fail on empty list")
	}
}
