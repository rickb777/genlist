package main

import (
	"testing"
	"reflect"
)

func TestMinByNum1(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 3},
		Thing{"Fie", 2},
		Thing{"Foe", 1},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	number := func(x Thing) Num1 {
		return x.Number
	}

	min1, err := things.MinByNum1(number)
	expected1 := Thing{"Foe", 1}

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
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	number := func(x Thing) Foo {
		return Foo(x.Name)
	}

	min1, err := things.MinByFoo(number)
	expected1 := Thing{"Bam", 13}

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
