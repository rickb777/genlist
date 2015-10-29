package main

import (
	"reflect"
	"testing"
)

func TestFlatMapToNum(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number0 := func(x Thing) Num1Collection {
		return num1Collection()
	}

	r0 := things.FlatMapToNum1(number0)
	expected0 := num1Collection()

	if !reflect.DeepEqual(r0, expected0) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %d %#v", expected0, r0.Size(), r0)
	}

	number1 := func(x Thing) Num1Collection {
		return num1Collection(x.Number, x.Number)
	}

	r1 := things.FlatMapToNum1(number1)
	expected1 := num1Collection(60, 60, -20, -20, 100, 100)

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %#v", expected1, r1)
	}

	number2 := func(x Thing) Num2Collection {
		v := Num2(x.Number)
		return num2Collection(&v, &v)
	}

	r2 := things.FlatMapToNum2(number2)
	expected2 := num2Collection(ip(60), ip(60), ip(-20), ip(-20), ip(100), ip(100))

	if !reflect.DeepEqual(r2, expected2) {
		t.Errorf("FlatMapToNum2 should result in %#v, got %#v", expected2, r2)
	}
}

func TestFlatMapToString(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	fn0 := func(x Thing) []string {
		return []string{}
	}

	r0 := things.FlatMapToString(fn0)
	expected0 := []string{}

	if !reflect.DeepEqual(r0, expected0) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %d %#v", expected0, len(r0), r0)
	}

	fn1 := func(x Thing) []string {
		return []string{x.Name, x.Name}
	}

	r1 := things.FlatMapToString(fn1)
	expected1 := []string{"First", "First", "Second", "Second", "Third", "Third"}

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %#v", expected1, r1)
	}
}
