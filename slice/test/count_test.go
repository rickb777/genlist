package main

import "testing"

func TestCountOther(t *testing.T) {
	things := OtherSlice{50, 100, 9, 7, 100, 99}

	count1 := things.Count(7)

	if count1 != 1 {
		t.Errorf("Count should find one 7")
	}

	count2 := things.Count(3)

	if count2 != 0 {
		t.Errorf("Count should no items for 3")
	}

	count3 := OtherSlice{}.Count(9)

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty slice")
	}
}

func TestCountThing(t *testing.T) {
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
