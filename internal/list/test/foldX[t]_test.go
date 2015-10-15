package main

import (
	"testing"
	"fmt"
)

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
		t.Errorf("FoldLeftNum1 should be %v, got %v", expected1, aggregate1)
	}
}

func TestFoldLeftColour(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Colour, x Thing) Colour {
		return Colour(fmt.Sprintf("%s:%d", state, x.Number))
	}

	aggregate1 := things.FoldLeftColour(Colour("red"), sum)
	expected1 := Colour("red:60:-20:100")

	if aggregate1 != expected1 {
		t.Errorf("FoldLeftColour should be %v, got %v", expected1, aggregate1)
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
		t.Errorf("FoldRightNum1 should be %v, got %v", expected1, aggregate1)
	}
}

func TestFoldRightColour(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Colour, x Thing) Colour {
		return Colour(fmt.Sprintf("%s:%d", state, x.Number))
	}

	aggregate1 := things.FoldRightColour(Colour("red"), sum)
	expected1 := Colour("red:100:-20:60")

	if aggregate1 != expected1 {
		t.Errorf("FoldRightColour should be %v, got %v", expected1, aggregate1)
	}
}
