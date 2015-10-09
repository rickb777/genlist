package main

import "testing"

func TestMeanOther(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -10},
		{"Third", 100},
	}

	number := func(x Thing) Other {
		return x.Number
	}

	mean1, err := things.MeanOther(number)

	if err != nil {
		t.Errorf("Mean should succeed")
	}

	expected1 := Other(50)

	if mean1 != expected1 {
		t.Errorf("Mean should be %v, got %v", expected1, mean1)
	}

	mean2, err := ThingList{}.MeanOther(number)

	if err == nil || mean2 != 0 {
		t.Errorf("Mean should fail on empty slice")
	}
}
