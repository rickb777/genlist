package main

import (
	"reflect"
	"testing"
)

func TestDistinctNum1(t *testing.T) {
	things := Num1List{50, 100, 9, 7, 100, 99}

	should := Num1List{50, 100, 9, 7, 99}

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %#v, but found %#v", should, distinct1)
	}
}

func TestDistinctNum2(t *testing.T) {
	things := Num2List{ip(50), ip(100), ip(9), ip(7), ip(100), ip(99)}

	should := Num2List{ip(50), ip(100), ip(9), ip(7), ip(99)}

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %#v, but found %#v", should, distinct1)
	}
}

func TestDistinctThing(t *testing.T) {
	things := ThingList{
		{"First", 0},
		{"Second", 0},
		{"First", 0},
		{"Third", 0},
	}

	should := ThingList{
		{"First", 0},
		{"Second", 0},
		{"Third", 0},
	}

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %#v, but found %#v", should, distinct1)
	}
}
