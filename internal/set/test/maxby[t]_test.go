package main

import (
	"testing"
	"reflect"
)

func TestMaxByNum1(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 21},
		Thing{"Bam", 13},
		Thing{"Bam", 8},
	)

	number := func(x Thing) Num1 {
		return x.Number
	}

	max1 := things.MaxByNum1(number)
	expected1 := Thing{"Boo", 21}

	if !reflect.DeepEqual(max1, expected1)  {
		t.Errorf("MaxByNum1 should be %#v, got %#v", expected1, max1)
	}
}

func TestMaxByFoo(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 3},
		Thing{"Fie", 2},
		Thing{"Foe", 1},
		Thing{"Boo", 5},
		Thing{"Boo", 21},
		Thing{"Bam", 13},
		Thing{"Bam", 8},
	)

	p := func(x Thing) Foo {
		return Foo(x.Name)
	}

	max1 := things.MaxByFoo(p)
	expected1 := Thing{"Foe", 1}

	if !reflect.DeepEqual(max1, expected1)  {
		t.Errorf("MaxByFoo should be %#v, got %#v", expected1, max1)
	}
}
