package main

import "testing"

func TestCountNum1(t *testing.T) {
	things := num1Collection(50, 100, 9, 7, 100, 99)

	count1 := things.Count(7)

	if count1 != 1 {
		t.Errorf("Count should find one 7")
	}

	count2 := things.Count(3)

	if count2 != 0 {
		t.Errorf("Count should no items for 3")
	}

	count3 := num1Collection().Count(9)

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty list")
	}
}

func TestCountNum2(t *testing.T) {
	things := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(100), ip(99))

	count1 := things.Count(ip(7))

	if count1 != 1 {
		t.Errorf("Count should find one 7")
	}

	count2 := things.Count(ip(3))

	if count2 != 0 {
		t.Errorf("Count should no items for 3")
	}

	count3 := num2Collection().Count(ip(9))

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty list")
	}
}

func TestCountThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	count1 := things.Count(Thing{"Fie", 2})

	if count1 != 1 {
		t.Errorf("Count should find one item Name == Second")
	}

	count2 := things.Count(Thing{"Dummy", 111})

	if count2 != 0 {
		t.Errorf("Count should no items for Name == Dummy")
	}

	count3 := ThingList{}.Count(Thing{"Fie", 200})

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty list")
	}
}
