package main

import (
	"reflect"
	"testing"
)

func TestMapToNum(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number1 := func(x Thing) Num1 {
		return x.Number
	}

	r1 := things.MapToNum1(number1)
	expected1 := num1Collection(60, -20, 100)

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("MapToNum1 should result in %#v, got %#v", expected1, r1)
	}

	number2 := func(x Thing) *Num2 {
		v := Num2(x.Number)
		return &v
	}

	r2 := things.MapToNum2(number2)
	expected2 := num2Collection(ip(60), ip(-20), ip(100))

	if !reflect.DeepEqual(r2, expected2) {
		t.Errorf("MapToNum2 should result in %#v, got %#v", expected2, r2)
	}
}

func TestMapToNumEmpty(t *testing.T) {
	noThings := ThingList{}

	number1 := func(x Thing) Num1 {
		return x.Number
	}

	r1 := noThings.MapToNum1(number1)
	expected1 := num1Collection()

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("MapToNum1 should result in %#v, got %#v", expected1, r1)
	}
}
