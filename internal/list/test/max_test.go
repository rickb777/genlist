package main

import "testing"

func TestMaxOrdered(t *testing.T) {
	others := num1Collection(50, 100, 9, 7, 100, 99)

	max1, err := others.Max()
	m1 := Num1(100)

	if err != nil {
		t.Errorf("Max should succeed")
	}

	if max1 != m1 {
		t.Errorf("Max should return %#v, got %#v", m1, max1)
	}

	max2, err := num1Collection().Max()
	var m2 Num1

	if err == nil || max2 != m2 {
		t.Errorf("Max should fail on empty list")
	}
}

func TestMaxNotOrdered(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
		{"Fourth", 10},
	}

	max1, err1 := things.Max(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	if err1 != nil {
		t.Errorf("Max Number should succeed")
	}

	expected1 := Thing{"Third", 100}

	if max1 != expected1 {
		t.Errorf("Max Number should return %#v, got %#v", expected1, max1)
	}

	_, err2 := ThingList{}.Max(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("Max Number should fail on empty list")
	}
}
