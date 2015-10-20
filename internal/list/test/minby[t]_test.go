package main

import (
	"testing"
	"reflect"
)

func TestMinByNum1(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number := func(x Thing) Num1 {
		return x.Number
	}

	min1, err := things.MinByNum1(number)
	expected1 := Thing{"Second", -20}

	if err != nil {
		t.Errorf("MinNum should succeed")
	}

	if !reflect.DeepEqual(min1, expected1)  {
		t.Errorf("MaxByNum1 should be %#v, got %#v", expected1, min1)
	}

	_, err = ThingList{}.MinByNum1(number)

	if err == nil {
		t.Errorf("MinNum should fail on empty list")
	}
}

func TestMinByColour(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
		{"Fourth", 1},
		{"Fifth", 17},
	}

	number := func(x Thing) Colour {
		return Colour(x.Name)
	}

	min1, err := things.MinByColour(number)
	expected1 := Thing{"Fifth", 17}

	if err != nil {
		t.Errorf("MinColour should succeed")
	}

	if !reflect.DeepEqual(min1, expected1)  {
		t.Errorf("MinColour should be %#v, got %#v", expected1, min1)
	}

	_, err = ThingList{}.MinByColour(number)

	if err == nil {
		t.Errorf("MinColour should fail on empty list")
	}
}
