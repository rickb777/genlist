package main

import "testing"

func TestMaxOrdered(t *testing.T) {
	others := OtherSlice{50, 100, 9, 7, 100, 99}

	max1, err := others.Max()
	m1 := Other(100)

	if err != nil {
		t.Errorf("Max should succeed")
	}

	if max1 != m1 {
		t.Errorf("Max should return %v, got %v", m1, max1)
	}

	max2, err := OtherSlice{}.Max()
	var m2 Other

	if err == nil || max2 != m2 {
		t.Errorf("Max should fail on empty slice")
	}
}

func TestMaxNotOrdered(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
		{"Fourth", 10},
	}

	max1, err1 := things.Max(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	if err1 != nil {
		t.Errorf("MaxBy Number should succeed")
	}

	expected1 := Thing{"Third", 100}

	if max1 != expected1 {
		t.Errorf("MaxBy Number should return %v, got %v", expected1, max1)
	}

	_, err2 := ThingSlice{}.Max(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MaxBy Number should fail on empty slice")
	}
}
