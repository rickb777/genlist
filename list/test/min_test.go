package main

import "testing"

func TestMinOrdered(t *testing.T) {
	others := OtherList{50, 100, 9, 7, 100, 99}

	min1, err := others.Min()
	m1 := Other(7)

	if err != nil {
		t.Errorf("Min should succeed")
	}

	if min1 != m1 {
		t.Errorf("Min should return %v, got %v", m1, min1)
	}

	min2, err := OtherList{}.Min()
	var m2 Other

	if err == nil || min2 != m2 {
		t.Errorf("Min should fail on empty slice")
	}
}

func TestMinNotOrdered(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
		{"Fourth", 10},
	}

	max1, err1 := things.Min(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	if err1 != nil {
		t.Errorf("Min Number should succeed")
	}

	expected1 := Thing{"Second", -20}

	if max1 != expected1 {
		t.Errorf("Min Number should return %v, got %v", expected1, max1)
	}

	_, err2 := ThingList{}.Min(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("Min Number should fail on empty slice")
	}
}
