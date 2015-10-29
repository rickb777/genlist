package main

import (
	"reflect"
	"testing"
)

func TestFilterNum1(t *testing.T) {
	things := num1Collection(60, 20, 100, 20)

	where1 := things.Filter(func(x Num1) bool {
		return x == 20
	})

	expected1 := num1Collection(20, 20)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Filter should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Filter(func(x Num1) bool {
		return x == 1
	})

	if where2.Size() != 0 {
		t.Errorf("Filter should result in empty list, got %#v", where2)
	}

	where3 := num1Collection().Filter(func(x Num1) bool {
		return true
	})

	if where3.Size() != 0 {
		t.Errorf("Filter should result in empty list, got %#v", where3)
	}
}

func TestFilterThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.Filter(func(x Thing) bool {
		return x.Name == "Boo"
	})

	expected1 := thingCollection(
		Thing{"Boo", 5},
		Thing{"Boo", 8},
	)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Filter should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Filter(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if where2.Size() != 0 {
		t.Errorf("Filter should result in empty list, got %#v", where2)
	}

	where3 := thingCollection().Filter(func(x Thing) bool {
		return true
	})

	if where3.Size() != 0 {
		t.Errorf("Filter should result in empty list, got %#v", where3)
	}
}
