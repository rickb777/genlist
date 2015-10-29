package main

import (
	"reflect"
	"testing"
)

func TestFlatMapToNum(t *testing.T) {
	things := thingCollection(
		Thing{"First", 60},
		Thing{"Second", -20},
		Thing{"Third", 100},
	)

	number0 := func(x Thing) Num1Collection {
		return num1Collection()
	}

	r0 := things.FlatMapToNum1(number0)
	expected0 := num1Collection()

	if !reflect.DeepEqual(r0, expected0) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %d %#v", expected0, r0.Size(), r0)
	}

	number1 := func(x Thing) Num1Collection {
		return num1Collection(x.Number, x.Number)
	}

	r1 := things.FlatMapToNum1(number1)
	expected1 := num1Collection(60, 60, -20, -20, 100, 100)

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %#v", expected1, r1)
	}

	number2 := func(x Thing) Num2Collection {
		v := Num2(x.Number)
		return num2Collection(&v, &v)
	}

	r2 := things.FlatMapToNum2(number2)
	expected2 := num2Collection(ip(60), ip(60), ip(-20), ip(-20), ip(100), ip(100))

	if !reflect.DeepEqual(r2, expected2) {
		t.Errorf("FlatMapToNum2 should result in %#v, got %#v", expected2, r2)
	}
}

func TestFlatMapToString(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	fn0 := func(x Thing) []string {
		return []string{}
	}

	r0 := things.FlatMapToString(fn0)
	expected0 := []string{}

	if !reflect.DeepEqual(r0, expected0) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %d %#v", expected0, len(r0), r0)
	}

	fn1 := func(x Thing) []string {
		return []string{x.Name, x.Name}
	}

	r1 := things.FlatMapToString(fn1)
	expected1 := []string{"Fee", "Fee", "Fie", "Fie", "Foe", "Foe", "Boo", "Boo", "Boo", "Boo", "Bam", "Bam", "Bam", "Bam"}

	if !reflect.DeepEqual(r1, expected1) {
		t.Errorf("FlatMapToNum1 should result in %#v, got %#v", expected1, r1)
	}
}
