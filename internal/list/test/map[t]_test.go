package main

import (
	"reflect"
	"testing"
)

func TestSelectOther(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number := func(x Thing) Other {
		return x.Number
	}

	select1 := things.MapToOther(number)
	expected1 := OtherList{60, -20, 100}

	if !reflect.DeepEqual(select1, expected1) {
		t.Errorf("MapToOther should result in %v, got %v", expected1, select1)
	}
}
