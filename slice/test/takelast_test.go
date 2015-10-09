package main

import (
	"reflect"
	"testing"
)

func TestTakeLastOther(t *testing.T) {
	things := OtherSlice{1, 3, 13, 17, 5, 6, 17, 8, 9}

	where1 := things.TakeLast(4)

	expected1 := OtherSlice{6, 17, 8, 9}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("TakeLast should result in %v, got %v", expected1, where1)
	}

	where2 := things.TakeLast(0)

	if len(where2) != 0 {
		t.Errorf("TakeLast should result in empty slice, got %v", where2)
	}

	where3 := things.TakeLast(100)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("TakeLast should result in %v, got %v", things, where3)
	}

	where4 := OtherSlice{}.TakeLast(100)

	if len(where4) != 0 {
		t.Errorf("TakeLast should result in empty slice, got %v", where4)
	}
}

func TestTakeLastThing(t *testing.T) {
	things := ThingSlice{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
		{"Boo", 4},
		{"Bam", 5},
	}

	where1 := things.TakeLast(2)

	expected1 := ThingSlice{
		{"Boo", 4},
		{"Bam", 5},
	}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("TakeLast should result in %v, got %v", expected1, where1)
	}

	where2 := things.TakeLast(0)

	if len(where2) != 0 {
		t.Errorf("TakeLast should result in empty slice, got %v", where2)
	}

	where3 := things.TakeLast(5)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("TakeLast should result in %v, got %v", things, where3)
	}

	where4 := ThingSlice{}.TakeLast(100)

	if len(where4) != 0 {
		t.Errorf("TakeLast should result in empty slice, got %v", where4)
	}
}
