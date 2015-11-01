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

	min1 := things.MinByNum1(number)
	expected1 := Thing{"Foe", 1}

	if !reflect.DeepEqual(min1, expected1)  {
		t.Errorf("MaxByNum1 should be %#v, got %#v", expected1, min1)
	}
}

func TestMinByFoo(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Bam", 13},
	)

	number := func(x Thing) Foo {
		return Foo(x.Name)
	}

	min1 := things.MinByFoo(number)
	expected1 := Thing{"Bam", 13}

	if !reflect.DeepEqual(min1, expected1)  {
		t.Errorf("MinFoo should be %#v, got %#v", expected1, min1)
	}
}
