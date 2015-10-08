package main

import "testing"

func TestMaxByOther(t *testing.T) {
	things := OtherSlice{50, 100, -20, 7, 100, 99}

	min1, err1 := things.MaxBy(func(a, b Other) bool {
		return a < b
	})

	if err1 != nil {
		t.Errorf("MaxBy Number should succeed")
	}

	if min1 != 100 {
		t.Errorf("MaxBy Number should return %v, got %v", 100, min1)
	}

	_, err2 := OtherSlice{}.MaxBy(func(a, b Other) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MaxBy Number should fail on empty slice")
	}
}

func TestMaxByThing(t *testing.T) {
	things := ThingSlice{
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

	_, err2 := ThingSlice{}.MaxBy(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MaxBy Number should fail on empty slice")
	}
}
