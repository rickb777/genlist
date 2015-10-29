package main

import (
	"reflect"
	"testing"
)

func TestDistinctNum1(t *testing.T) {
	things := num1Collection(50, 100, 9, 7, 100, 99)

	should := num1Collection(50, 100, 9, 7, 99)

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %#v, but found %#v", should, distinct1)
	}
}

func TestDistinctNum2(t *testing.T) {
	things := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(100), ip(99))

	should := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(99))

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %#v, but found %#v", should, distinct1)
	}
}

func TestDistinctThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 0},
		Thing{"Fie", 0},
		Thing{"Foe", 0},
		Thing{"Boo", 0},
		Thing{"Boo", 0},
		Thing{"Bam", 0},
		Thing{"Bam", 0},
	)

	should := thingCollection(
		Thing{"Fee", 0},
		Thing{"Fie", 0},
		Thing{"Foe", 0},
		Thing{"Boo", 0},
		Thing{"Bam", 0},
	)

	distinct1 := things.Distinct()

	if !reflect.DeepEqual(distinct1, should) {
		t.Errorf("Distinct should exclude be %#v, but found %#v", should, distinct1)
	}
}
