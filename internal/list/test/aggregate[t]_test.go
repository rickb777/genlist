package main

import "testing"

func TestAggregateNum1(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Num1, x Thing) Num1 {
		return state + x.Number
	}

	aggregate1 := things.AggregateNum1(sum)
	expected1 := Num1(140)

	if aggregate1 != expected1 {
		t.Errorf("AggregateNum1 should be %v, got %v", expected1, aggregate1)
	}
}
