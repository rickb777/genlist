package main

import (
	"reflect"
	"testing"
)

func TestDropLastOther(t *testing.T) {
	things := OtherSlice{1, 17, 5, 6, 17, 8, 9}

	where1 := things.DropLast(3)

	expected1 := OtherSlice{1, 17, 5, 6}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("DropLast should result in %v, got %v", expected1, where1)
	}

	where2 := things.DropLast(0)

	if !reflect.DeepEqual(where2, things) {
		t.Errorf("DropLast should result in %v, got %v", things, where2)
	}

	where3 := things.DropLast(100)

	if len(where3) != 0 {
		t.Errorf("DropLast should result in empty slice, got %v", where3)
	}

	where4 := OtherSlice{}.DropLast(100)

	if len(where4) != 0 {
		t.Errorf("DropLast should result in empty slice, got %v", where4)
	}
}

func TestDropLastThing(t *testing.T) {
	things := ThingSlice{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
		{"Boo", 4},
		{"Bam", 5},
	}

	where1 := things.DropLast(2)

	expected1 := ThingSlice{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
	}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("DropLast should result in %v, got %v", expected1, where1)
	}

	where2 := things.DropLast(10)

	if len(where2) != 0 {
		t.Errorf("DropLast should result in empty slice, got %v", where2)
	}

	where3 := things.DropLast(0)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("DropLast should result in %v, got %v", things, where3)
	}

	where4 := ThingSlice{}.DropLast(10)

	if len(where4) != 0 {
		t.Errorf("DropLast should result in empty slice, got %v", where4)
	}
}
