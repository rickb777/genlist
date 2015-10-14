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

	number := func(x Thing) Num1 {
		return x.Number
	}

	select1 := things.MapToNum1(number)
	expected1 := Num1List{60, -20, 100}

	if !reflect.DeepEqual(select1, expected1) {
		t.Errorf("MapToNum should result in %v, got %v", expected1, select1)
	}
}
