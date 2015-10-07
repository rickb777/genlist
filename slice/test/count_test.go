package main

import "testing"

func TestCount(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", 20},
		{"Third", 100},
	}

	count1 := things.Count(Thing{"Second", 20})

	if count1 != 1 {
		t.Errorf("Count should find one item Name == Second")
	}

	count2 := things.Count(Thing{"Dummy", 20})

	if count2 != 0 {
		t.Errorf("Count should no items for Name == Dummy")
	}

	count3 := ThingSlice{}.Count(Thing{"Second", 20})

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty slice")
	}
}
