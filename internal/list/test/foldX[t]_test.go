package main

import "testing"

func TestFoldLeftNum1(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Num1, x Thing) Num1 {
		return state + x.Number
	}

	aggregate1 := things.FoldLeftNum1(5, sum)
	expected1 := Num1(145)

	if aggregate1 != expected1 {
		t.Errorf("AggregateNum1 should be %v, got %v", expected1, aggregate1)
	}
}

func TestFoldRightNum1(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Num1, x Thing) Num1 {
		return state + x.Number
	}

	aggregate1 := things.FoldRightNum1(5, sum)
	expected1 := Num1(145)

	if aggregate1 != expected1 {
		t.Errorf("AggregateNum1 should be %v, got %v", expected1, aggregate1)
	}
}
