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
	things := ThingList{
		{"First", 0},
		{"Second", 9},
		{"First", 4},
		{"Third", 9},
		{"Fourth", 5},
		{"Fifth", 4},
	}

	expected := ThingList{
		{"First", 0},
		{"Second", 9},
		{"First", 4},
		{"Fourth", 5},
	}

	distinctby1 := things.DistinctBy(func(a, b Thing) bool {
		return a.Number == b.Number
	})

	if !reflect.DeepEqual(distinctby1, expected) {
		t.Errorf("DistinctBy should be %#v, but got %#v", expected, distinctby1)
	}
}
