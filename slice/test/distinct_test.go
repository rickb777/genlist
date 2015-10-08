package main

import (
	"reflect"
	"testing"
)

func TestDistinctOther(t *testing.T) {
	things := OtherSlice{50, 100, 9, 7, 100, 99}

	should := OtherSlice{50, 100, 9, 7, 99}

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %v, but found %v", should, distinct1)
	}
}

func TestDistinctThing(t *testing.T) {
	things := ThingSlice{
		{"First", 0},
		{"Second", 0},
		{"First", 0},
		{"Third", 0},
	}

	should := ThingSlice{
		{"First", 0},
		{"Second", 0},
		{"Third", 0},
	}

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %v, but found %v", should, distinct1)
	}
}
