package main

import (
	"reflect"
	"testing"
)

func TestDropNum1(t *testing.T) {
	things := num1Collection(1, 17, 5, 6, 17, 8, 9)

	where1 := things.Drop(3)

	expected1 := num1Collection(6, 17, 8, 9)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Drop should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Drop(0)

	if !reflect.DeepEqual(where2, things) {
		t.Errorf("Drop should result in %#v, got %#v", things, where2)
	}

	where3 := things.Drop(100)

	if len(where3) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where3)
	}

	where4 := num1Collection().Drop(100)

	if len(where4) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where4)
	}
}

func TestDropNum2(t *testing.T) {
	things := num2Collection(ip(1), ip(17), ip(5), ip(6), ip(17), ip(8), ip(9))

	where1 := things.Drop(3)

	expected1 := num2Collection(ip(6), ip(17), ip(8), ip(9))

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Drop should result in %#v, got %#v", expected1, where1)
	}
}

func TestDropThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.Drop(4)

	expected1 := thingCollection(
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)


	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Drop should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Drop(10)

	if len(where2) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where2)
	}

	where3 := things.Drop(0)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("Drop should result in %#v, got %#v", things, where3)
	}

	where4 := thingCollection().Drop(10)

	if len(where4) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where4)
	}
}
