package main

import (
	"testing"
	"reflect"
)

func TestReverseNum(t *testing.T) {
	others := Num1List{50, 100, 9, 7, 100, 99}
	expected := Num1List{99, 100, 7, 9, 100, 50}

	rev1 := others.Reverse()

	if !reflect.DeepEqual(rev1, expected) {
		t.Errorf("Reverse should return %#v, got %#v", expected, rev1)
	}

	rev2 := Num1List{}.Reverse()

	if len(rev2) != 0 {
		t.Errorf("Reverse should handle empty list")
	}
}

func TestReverseThing(t *testing.T) {
	things := ThingList{
		{"First", 0},
		{"Second", 0},
		{"Third", 0},
		{"Second", 10},
	}
	expected := ThingList{
		{"Second", 10},
		{"Third", 0},
		{"Second", 0},
		{"First", 0},
	}

	rev1 := things.Reverse()

	if !reflect.DeepEqual(rev1, expected) {
		t.Errorf("Reverse should return %#v, got %#v", expected, rev1)
	}

	rev2 := Num1List{}.Reverse()

	if len(rev2) != 0 {
		t.Errorf("Reverse should handle empty list")
	}
}
