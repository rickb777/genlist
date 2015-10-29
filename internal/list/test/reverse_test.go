package main

import (
	"testing"
	"reflect"
)

func TestReverseNum(t *testing.T) {
	others := num1Collection(50, 100, 9, 7, 100, 99)
	expected := num1Collection(99, 100, 7, 9, 100, 50)

	rev1 := others.Reverse()

	if !reflect.DeepEqual(rev1, expected) {
		t.Errorf("Reverse should return %#v, got %#v", expected, rev1)
	}

	rev2 := num1Collection().Reverse()

	if len(rev2) != 0 {
		t.Errorf("Reverse should handle empty list")
	}
}

func TestReverseThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)
	expected := thingCollection(
		Thing{"Bam", 21},
		Thing{"Bam", 13},
		Thing{"Boo", 8},
		Thing{"Boo", 5},
		Thing{"Foe", 3},
		Thing{"Fie", 2},
		Thing{"Fee", 1},
	)

	rev1 := things.Reverse()

	if !reflect.DeepEqual(rev1, expected) {
		t.Errorf("Reverse should return %#v, got %#v", expected, rev1)
	}

	rev2 := num1Collection().Reverse()

	if len(rev2) != 0 {
		t.Errorf("Reverse should handle empty list")
	}
}
