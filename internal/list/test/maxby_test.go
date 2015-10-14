package main

import "testing"

func TestMaxByNum(t *testing.T) {
	things := Num1List{50, 100, -20, 7, 100, 99}

	min1, err1 := things.MaxBy(func(a, b Num1) bool {
		return a < b
	})

	if err1 != nil {
		t.Errorf("MaxBy Number should succeed")
	}

	if min1 != 100 {
		t.Errorf("MaxBy Number should return %v, got %v", 100, min1)
	}

	_, err2 := Num1List{}.MaxBy(func(a, b Num1) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MaxBy Number should fail on empty list")
	}
}

func TestMaxByThing(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	min1, err1 := things.MaxBy(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	if err1 != nil {
		t.Errorf("MaxBy Number should succeed")
	}

	expected1 := Thing{"Third", 100}
	if min1 != expected1 {
		t.Errorf("MaxBy Number should return %v, got %v", expected1, min1)
	}

	_, err2 := ThingList{}.MaxBy(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MaxBy Number should fail on empty list")
	}
}
