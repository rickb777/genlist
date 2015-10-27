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

func TestMinByFoo(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
		{"Fourth", 1},
		{"Fifth", 17},
	}

	number := func(x Thing) Foo {
		return Foo(x.Name)
	}

	min1, err := things.MinByFoo(number)
	expected1 := Thing{"Fifth", 17}

	if err != nil {
		t.Errorf("MinFoo should succeed")
	}

	if !reflect.DeepEqual(min1, expected1)  {
		t.Errorf("MinFoo should be %#v, got %#v", expected1, min1)
	}

	_, err = ThingList{}.MinByFoo(number)

	if err == nil {
		t.Errorf("MinFoo should fail on empty list")
	}
}
