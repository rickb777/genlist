package main

import (
	"reflect"
	"testing"
)

func TestGroupByNum1(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -10},
		{"Third", 100},
		{"Fourth", -10},
		{"Fifth", 60},
	}

	number := func(x Thing) Num1 {
		return x.Number
	}

	groupby1 := things.GroupByNum1(number)
	expected1 := map[Num1]ThingList{
		-10: {{"Second", -10}, {"Fourth", -10}},
		60:  {{"First", 60}, {"Fifth", 60}},
		100: {{"Third", 100}},
	}

	if len(groupby1) != len(expected1) {
		t.Errorf("GroupByInt result should have %d elements, has %d", len(expected1), len(groupby1))
	}

	if !reflect.DeepEqual(groupby1, expected1) {
		t.Errorf("GroupByNum should be %v, got %v", expected1, groupby1)
	}
}

func TestGroupByColour(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -10},
		{"Third", 100},
		{"Fourth", 0},
		{"Fifth", 60},
	}

	number := func(x Thing) Colour {
		switch {
		case x.Number < 10: return Colour("green")
		case x.Number < 70: return Colour("amber")
		default: return Colour("red")
		}
	}

	groupby1 := things.GroupByColour(number)
	expected1 := map[Colour]ThingList{
		Colour("green"): {{"Second", -10}, {"Fourth", 0}},
		Colour("amber"): {{"First", 60}, {"Fifth", 60}},
		Colour("red"): {{"Third", 100}},
	}

	if len(groupby1) != len(expected1) {
		t.Errorf("GroupByInt result should have %d elements, has %d", len(expected1), len(groupby1))
	}

	if !reflect.DeepEqual(groupby1, expected1) {
		t.Errorf("GroupByNum should be %v, got %v", expected1, groupby1)
	}
}
