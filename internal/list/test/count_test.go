package main

import "testing"

func ip(n int) *Num2 {
	v := Num2(n)
	return &v
}

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
	things := Num2List{ip(50), ip(100), ip(9), ip(7), ip(100), ip(99)}

	count1 := things.Count(ip(7))

	if count1 != 1 {
		t.Errorf("Count should find one 7")
	}

	count2 := things.Count(ip(3))

	if count2 != 0 {
		t.Errorf("Count should no items for 3")
	}

	count3 := Num2List{}.Count(ip(9))

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty list")
	}
}

func TestCountThing(t *testing.T) {
	things := ThingList{
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

	count3 := ThingList{}.Count(Thing{"Second", 20})

	if count3 != 0 {
		t.Errorf("Count should find no items in an empty list")
	}
}
