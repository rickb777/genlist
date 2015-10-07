package main

import (
	"testing"
	"reflect"
)

func TestReverse(t *testing.T) {
	others := OtherSlice{50, 100, 9, 7, 100, 99}
	expected := OtherSlice{99, 100, 7, 9, 100, 50}

	rev1 := others.Reverse()

	if !reflect.DeepEqual(rev1, expected) {
		t.Errorf("Reverse should return %v, got %v", expected, rev1)
	}

	rev2 := OtherSlice{}.Reverse()

	if len(rev2) != 0 {
		t.Errorf("Reverse should handle empty slice")
	}
}
