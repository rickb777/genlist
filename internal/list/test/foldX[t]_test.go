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
	things := thingCollection(
		Thing{"First", 60},
		Thing{"Second", -20},
		Thing{"Third", 100},
	)

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
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	sum := func(state Foo, x Thing) Foo {
		return Foo(fmt.Sprintf("%s:%d", state, x.Number))
	}

	aggregate1 := things.FoldRightFoo(Foo("red"), sum)
	expected1 := Foo("red:21:13:8:5:3:2:1")

	if aggregate1 != expected1 {
		t.Errorf("FoldRightFoo should be %#v, got %#v", expected1, aggregate1)
	}
}
