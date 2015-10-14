package main

import "testing"

func TestMaxNum(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number := func(x Thing) Num1 {
		return x.Number
	}

	max1, err := things.MaxNum1(number)

	if err != nil {
		t.Errorf("MaxNum should succeed")
	}

	if max1 != 100 {
		t.Errorf("MaxNum should be %v, got %v", 100, max1)
	}

	max2, err := ThingList{}.MaxNum1(number)

	if err == nil || max2 != 0 {
		t.Errorf("Max should fail on empty list")
	}
}
