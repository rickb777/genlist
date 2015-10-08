package main

import (
	"reflect"
	"testing"
)

func TestDistinctByOther(t *testing.T) {
	things := OtherSlice{50, 100, 9, 7, 9, 100, 99, 50}

	expected := OtherSlice{50, 100, 9, 7, 99}

	distinctby1 := things.DistinctBy(func(a, b Other) bool {
		return a == b
	})

	if !reflect.DeepEqual(distinctby1, expected) {
		t.Errorf("DistinctBy should be %v, but got %v", expected, distinctby1)
	}
}

func TestDistinctByThing(t *testing.T) {
	things := ThingSlice{
		{"First", 0},
		{"Second", 9},
		{"First", 4},
		{"Third", 9},
		{"Fourth", 5},
		{"Fifth", 4},
	}

	expected := ThingSlice{
		{"First", 0},
		{"Second", 9},
		{"First", 4},
		{"Fourth", 5},
	}

	distinctby1 := things.DistinctBy(func(a, b Thing) bool {
		return a.Number == b.Number
	})

	if !reflect.DeepEqual(distinctby1, expected) {
		t.Errorf("DistinctBy should be %v, but got %v", expected, distinctby1)
	}
}
