package main

import "testing"

func TestCountBy(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", 20},
		{"Third", 100},
	}

	count1 := things.CountBy(func(x Thing) bool {
		return x.Name == "Second"
	})

	if count1 != 1 {
		t.Errorf("CountBy should find one item Name == Second")
	}

	count2 := things.CountBy(func(x Thing) bool {
		return x.Number > 50
	})

	if count2 != 2 {
		t.Errorf("CountBy should find 2 items for Number > 50")
	}

	count3 := things.CountBy(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if count3 != 0 {
		t.Errorf("CountBy should no items for Name == Dummy")
	}

	count4 := ThingSlice{}.CountBy(func(x Thing) bool {
		return true
	})

	if count4 != 0 {
		t.Errorf("CountBy should find no items in an empty slice")
	}
}
