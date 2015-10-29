package main

import "testing"

func TestMeanNum(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	number := func(x Thing) Num1 {
		return x.Number
	}

	mean1, err := things.MeanNum1(number)

	if err != nil {
		t.Errorf("Mean should succeed")
	}

	expected1 := Num1(7)

	if mean1 != expected1 {
		t.Errorf("Mean should be %#v, got %#v", expected1, mean1)
	}

	mean2, err := ThingList{}.MeanNum1(number)

	if err == nil || mean2 != 0 {
		t.Errorf("Mean should fail on empty list")
	}
}
