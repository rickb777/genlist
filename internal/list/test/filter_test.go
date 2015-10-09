package main

import (
	"reflect"
	"testing"
)

func TestFilterOther(t *testing.T) {
	things := OtherList{60, 20, 100, 20}

	where1 := things.Filter(func(x Other) bool {
		return x == 20
	})

	expected1 := OtherList{20, 20}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Filter should result in %v, got %v", expected1, where1)
	}

	where2 := things.Filter(func(x Other) bool {
		return x == 1
	})

	if len(where2) != 0 {
		t.Errorf("Filter should result in empty list, got %v", where2)
	}

	where3 := OtherList{}.Filter(func(x Other) bool {
		return true
	})

	if len(where3) != 0 {
		t.Errorf("Filter should result in empty list, got %v", where3)
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
		t.Errorf("Filter should result in %v, got %v", expected1, where1)
	}

	where2 := things.Filter(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if len(where2) != 0 {
		t.Errorf("Filter should result in empty list, got %v", where2)
	}

	where3 := ThingList{}.Filter(func(x Thing) bool {
		return true
	})

	if len(where3) != 0 {
		t.Errorf("Filter should result in empty list, got %v", where3)
	}
}
