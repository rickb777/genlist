package main

import (
	"reflect"
	"testing"
)

func TestGroupByNum1(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 1},
		Thing{"Foe", 1},
		Thing{"Boo", 4},
		Thing{"Boo", 4},
		Thing{"Bam", 4},
		Thing{"Bam", 6},
	)

	number := func(x Thing) Num1 {
		return x.Number
	}

	groupby1 := things.GroupByNum1(number)
	expected1 := map[Num1]ThingList{
		1: {{"Fee", 1}, {"Fie", 1}, {"Foe", 1}},
		4: {{"Boo", 4}, {"Boo", 4}, {"Bam", 4}},
		6: {{"Bam", 6}},
	}

	if len(groupby1) != len(expected1) {
		t.Errorf("GroupByInt result should have %d elements, has %d", len(expected1), len(groupby1))
	}

	if !reflect.DeepEqual(groupby1, expected1) {
		t.Errorf("GroupByNum should be %#v, got %#v", expected1, groupby1)
	}
}

func TestGroupByFoo(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 20},
		Thing{"Foe", 3},
		Thing{"Boo", 40},
		Thing{"Boo", 5},
		Thing{"Bam", 60},
		Thing{"Bam", 7},
	)

	number := func(x Thing) Foo {
		switch {
		case x.Number <= 10: return Foo("green")
		case x.Number <= 50: return Foo("amber")
		default: return Foo("red")
		}
	}

	groupby1 := things.GroupByFoo(number)
	expected1 := map[Foo]ThingList{
		Foo("green"): {{"Fee", 1}, {"Foe", 3}, {"Boo", 5}, {"Bam", 7}},
		Foo("amber"): {{"Fie", 20}, {"Boo", 40}},
		Foo("red"): {{"Bam", 60}},
	}

	if len(groupby1) != len(expected1) {
		t.Errorf("GroupByInt result should have %d elements, has %d", len(expected1), len(groupby1))
	}

	if !reflect.DeepEqual(groupby1, expected1) {
		t.Errorf("GroupByNum should be %#v, got %#v", expected1, groupby1)
	}
}
