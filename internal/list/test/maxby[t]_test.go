package main

import (
	"testing"
	"reflect"
)

func TestMaxByNum1(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number := func(x Thing) Num1 {
		return x.Number
	}

	max1, err := things.MaxByNum1(number)
	expected1 := Thing{"Third", 100}

	if err != nil {
		t.Errorf("MaxByNum1 should succeed")
	}

	if !reflect.DeepEqual(max1, expected1)  {
		t.Errorf("MaxByNum1 should be %#v, got %#v", expected1, max1)
	}

	_, err = ThingList{}.MaxByNum1(number)

	if err == nil {
		t.Errorf("MaxByNum1 should fail on empty list")
	}
}

func TestMaxByFoo(t *testing.T) {
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

	max1, err := things.MaxByFoo(number)
	expected1 := Thing{"Third", 100}

	if err != nil {
		t.Errorf("MaxFoo should succeed")
	}

	if !reflect.DeepEqual(max1, expected1)  {
		t.Errorf("MaxFoo should be %#v, got %#v", expected1, max1)
	}

	_, err = ThingList{}.MaxByFoo(number)

	if err == nil {
		t.Errorf("MaxFoo should fail on empty list")
	}
}
