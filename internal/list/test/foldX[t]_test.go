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
		t.Errorf("FoldLeftNum1 should be %#v, got %#v", expected1, aggregate1)
	}
}

func TestFoldLeftFoo(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Foo, x Thing) Foo {
		return Foo(fmt.Sprintf("%s:%d", state, x.Number))
	}

	aggregate1 := things.FoldLeftFoo(Foo("red"), sum)
	expected1 := Foo("red:60:-20:100")

	if aggregate1 != expected1 {
		t.Errorf("FoldLeftFoo should be %#v, got %#v", expected1, aggregate1)
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
		t.Errorf("FoldRightNum1 should be %#v, got %#v", expected1, aggregate1)
	}
}

func TestFoldRightFoo(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	sum := func(state Foo, x Thing) Foo {
		return Foo(fmt.Sprintf("%s:%d", state, x.Number))
	}

	aggregate1 := things.FoldRightFoo(Foo("red"), sum)
	expected1 := Foo("red:100:-20:60")

	if aggregate1 != expected1 {
		t.Errorf("FoldRightFoo should be %#v, got %#v", expected1, aggregate1)
	}
}
