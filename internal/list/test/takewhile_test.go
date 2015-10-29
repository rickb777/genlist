package main

import (
	"reflect"
	"testing"
)

func TestTakeWhileNum(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9)

	where1 := things.TakeWhile(func(x Num1) bool {
		return x % 2 == 1
	})

	expected1 := num1Collection(1, 3, 17, 5)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("TakeWhile should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.TakeWhile(func(x Num1) bool {
		return x == 0
	})

	if len(where2) != 0 {
		t.Errorf("TakeWhile should result in empty list, got %#v", where2)
	}

	where3 := things.TakeWhile(func(x Num1) bool {
		return true
	})

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("TakeWhile should result in %#v, got %#v", things, where3)
	}

	where4 := num1Collection().TakeWhile(func(x Num1) bool {
		return true
	})

	if len(where4) != 0 {
		t.Errorf("TakeWhile should result in empty list, got %#v", where4)
	}
}

func TestTakeWhileThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Boo", 5},
		Thing{"Foe", 3},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.TakeWhile(func(x Thing) bool {
		return x.Name[0] == 'F'
	})

	expected1 := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
	)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("TakeWhile should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.TakeWhile(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if len(where2) != 0 {
		t.Errorf("TakeWhile should result in empty list, got %#v", where2)
	}

	where3 := things.TakeWhile(func(x Thing) bool {
		return true
	})

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("TakeWhile should result in %#v, got %#v", things, where3)
	}

	where4 := thingCollection().TakeWhile(func(x Thing) bool {
		return true
	})

	if len(where4) != 0 {
		t.Errorf("TakeWhile should result in empty list, got %#v", where4)
	}
}
