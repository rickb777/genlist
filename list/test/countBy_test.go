package main

import "testing"

func TestCountByOther(t *testing.T) {
	things := OtherList{50, 100, 9, 7, 100, 99}

	count1 := things.CountBy(func(x Other) bool {
		return x == 9
	})

	if count1 != 1 {
		t.Errorf("CountBy should find one item with the value 9")
	}

	count2 := things.CountBy(func(x Other) bool {
		return x > 50
	})

	if count2 != 3 {
		t.Errorf("CountBy should find 3 items > 50")
	}

	count3 := things.CountBy(func(x Other) bool {
		return x == 1
	})

	if count3 != 0 {
		t.Errorf("CountBy should no items with the value 1")
	}

	count4 := OtherList{}.CountBy(func(x Other) bool {
		return true
	})

	if count4 != 0 {
		t.Errorf("CountBy should find no items in an empty slice")
	}
}

func TestCountByThing(t *testing.T) {
	things := ThingList{
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

	count4 := ThingList{}.CountBy(func(x Thing) bool {
		return true
	})

	if count4 != 0 {
		t.Errorf("CountBy should find no items in an empty slice")
	}
}
