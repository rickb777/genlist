package main

import (
	"reflect"
	"testing"
)

func TestSelectNum(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number1 := func(x Thing) Num1 {
		return x.Number
	}

	select1 := things.MapToNum1(number1)
	expected1 := Num1List{60, -20, 100}

	if !reflect.DeepEqual(select1, expected1) {
		t.Errorf("MapToNum should result in %v, got %v", expected1, select1)
	}

	number2 := func(x Thing) *Num2 {
		v := Num2(x.Number)
		return &v
	}

	select2 := things.MapToNum2(number2)
	expected2 := Num2List{ip(60), ip(-20), ip(100)}

	if !reflect.DeepEqual(select2, expected2) {
		t.Errorf("MapToNum should result in %v, got %v", expected2, select2)
	}
}
