package main

import (
	"reflect"
	"testing"
)

func TestMapToNum(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	number1 := func(x Thing) Num1 {
		return x.Number
	}

	r1 := things.MapToNum1(number1)
	expected1 := num1Collection(1, 2, 3, 5, 8, 13, 21)

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("MapToNum1 should result in %#v, got %#v", expected1, r1)
	}
}

func TestMapToString(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Bam", 13},
	)

	name1 := func(x Thing) string {
		return x.Name
	}

	r1 := things.MapToString(name1)
	e1 := []string{"Fee", "Fie", "Foe", "Boo", "Bam"}

	if len(r1) != len(e1) || r1[0] == r1[1] {
		t.Errorf("MapToNum1 got %#v", r1)
	}
}

func TestMapToNumEmpty(t *testing.T) {
	noThings := thingCollection()

	number1 := func(x Thing) Num1 {
		return x.Number
	}

	r1 := noThings.MapToNum1(number1)
	expected1 := num1Collection()

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("MapToNum1 should result in %#v, got %#v", expected1, r1)
	}
}
