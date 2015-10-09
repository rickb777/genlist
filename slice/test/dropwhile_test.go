package main

import (
	"reflect"
	"testing"
)

func TestDropWhileOther(t *testing.T) {
	things := OtherSlice{1, 3, 17, 5, 6, 17, 8, 9}

	where1 := things.DropWhile(func(x Other) bool {
		return x % 2 == 1
	})

	expected1 := OtherSlice{6, 17, 8, 9}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("DropWhile should result in %v, got %v", expected1, where1)
	}

	where2 := things.DropWhile(func(x Other) bool {
		return x == 0
	})

	if !reflect.DeepEqual(where2, things) {
		t.Errorf("DropWhile should result in %v, got %v", things, where2)
	}

	where3 := things.DropWhile(func(x Other) bool {
		return true
	})

	if len(where3) != 0 {
		t.Errorf("DropWhile should result in empty slice, got %v", where3)
	}

	where4 := OtherSlice{}.DropWhile(func(x Other) bool {
		return true
	})

	if len(where4) != 0 {
		t.Errorf("DropWhile should result in empty slice, got %v", where4)
	}
}

func TestDropWhileThing(t *testing.T) {
	things := ThingSlice{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
		{"Boo", 4},
		{"Bam", 5},
	}

	where1 := things.DropWhile(func(x Thing) bool {
		return x.Name[0] == 'F'
	})

	expected1 := ThingSlice{
		{"Boo", 4},
		{"Bam", 5},
	}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("DropWhile should result in %v, got %v", expected1, where1)
	}

	where2 := things.DropWhile(func(x Thing) bool {
		return true
	})

	if len(where2) != 0 {
		t.Errorf("DropWhile should result in empty slice, got %v", where2)
	}

	where3 := things.DropWhile(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("DropWhile should result in %v, got %v", things, where3)
	}

	where4 := ThingSlice{}.DropWhile(func(x Thing) bool {
		return true
	})

	if len(where4) != 0 {
		t.Errorf("DropWhile should result in empty slice, got %v", where4)
	}
}
