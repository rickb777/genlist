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
	things := ThingList{
		{"First", 0},
		{"Second", 0},
		{"Third", 0},
		{"Second", 10},
	}

	where1 := things.Filter(func(x Thing) bool {
		return x.Name == "Second"
	})

	expected1 := ThingList{
		{"Second", 0},
		{"Second", 10},
	}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Filter should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Filter(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if where2.Size() != 0 {
		t.Errorf("Filter should result in empty list, got %#v", where2)
	}

	where3 := ThingList{}.Filter(func(x Thing) bool {
		return true
	})

	if where3.Size() != 0 {
		t.Errorf("Filter should result in empty list, got %#v", where3)
	}
}
