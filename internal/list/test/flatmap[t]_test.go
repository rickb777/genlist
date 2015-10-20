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

	number0 := func(x Thing) Num1Seq {
		return Num1List{}
	}

	r0 := things.FlatMapToNum1(number0)
	expected0 := Num1List{}

	if !reflect.DeepEqual(r0, expected0) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %d %#v", expected0, r0.Size(), r0)
	}

	number1 := func(x Thing) Num1Seq {
		return Num1List{x.Number, x.Number}
	}

	r1 := things.FlatMapToNum1(number1)
	expected1 := Num1List{60, 60, -20, -20, 100, 100}

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %#v", expected1, r1)
	}

	number2 := func(x Thing) Num2Seq {
		v := Num2(x.Number)
		return Num2List{&v, &v}
	}

	r2 := things.FlatMapToNum2(number2)
	expected2 := Num2List{ip(60), ip(60), ip(-20), ip(-20), ip(100), ip(100)}

	if !reflect.DeepEqual(r2, expected2) {
		t.Errorf("FlatMapToNum2 should result in %#v, got %#v", expected2, r2)
	}
}
