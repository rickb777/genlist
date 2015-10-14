package main

import "testing"

func TestMinNum(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number := func(x Thing) Num1 {
		return x.Number
	}

	min1, err := things.MinNum1(number)

	if err != nil {
		t.Errorf("MinNum should succeed")
	}

	if min1 != -20 {
		t.Errorf("MinNum should be %v, got %v", -20, min1)
	}

	min2, err := ThingList{}.MinNum1(number)

	if err == nil || min2 != 0 {
		t.Errorf("MinNum should fail on empty list")
	}
}
