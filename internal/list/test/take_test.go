package main

import (
	"reflect"
	"testing"
)

func TestTakeNum(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9, 13)

	where1 := things.Take(4)

	expected1 := num1Collection(1, 3, 17, 5)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Take should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Take(0)

	if len(where2) != 0 {
		t.Errorf("Take should result in empty list, got %#v", where2)
	}

	where3 := things.Take(100)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("Take should result in %#v, got %#v", things, where3)
	}

	where4 := num1Collection().Take(100)

	if len(where4) != 0 {
		t.Errorf("Take should result in empty list, got %#v", where4)
	}
}

func TestTakeThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.Take(3)

	expected1  := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
	)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Take should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Take(0)

	if len(where2) != 0 {
		t.Errorf("Take should result in empty list, got %#v", where2)
	}

	where3 := things.Take(100)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("Take should result in %#v, got %#v", things, where3)
	}

	where4 := thingCollection().Take(100)

	if len(where4) != 0 {
		t.Errorf("Take should result in empty list, got %#v", where4)
	}
}
