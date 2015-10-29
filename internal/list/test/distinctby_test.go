package main

import (
	"reflect"
	"testing"
)

func TestDistinctByNum1(t *testing.T) {
	things := num1Collection(50, 100, 9, 7, 9, 100, 99, 50)

	expected := num1Collection(50, 100, 9, 7, 99)

	distinctby1 := things.DistinctBy(func(a, b Num1) bool {
		return a == b
	})

	if !reflect.DeepEqual(distinctby1, expected) {
		t.Errorf("DistinctBy should be %#v, but got %#v", expected, distinctby1)
	}
}

func TestDistinctByNum2(t *testing.T) {
	things := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(100), ip(99))

	expected := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(99))

	distinctby1 := things.DistinctBy(func(a, b *Num2) bool {
		return *a == *b
	})

	if !reflect.DeepEqual(distinctby1, expected) {
		t.Errorf("DistinctBy should be %#v, but got %#v", expected, distinctby1)
	}
}

func TestDistinctByThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 4},
		Thing{"Boo", 4},
		Thing{"Bam", 4},
		Thing{"Bam", 4},
	)

	expected := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 4},
	)

	distinctby1 := things.DistinctBy(func(a, b Thing) bool {
		return a.Number == b.Number
	})

	if !reflect.DeepEqual(distinctby1, expected) {
		t.Errorf("DistinctBy should be %#v, but got %#v", expected, distinctby1)
	}
}
