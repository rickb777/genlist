package main

import (
	"reflect"
	"testing"
)

func TestDropNum1(t *testing.T) {
	things := Num1List{1, 17, 5, 6, 17, 8, 9}

	where1 := things.Drop(3)

	expected1 := Num1List{6, 17, 8, 9}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Drop should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Drop(0)

	if !reflect.DeepEqual(where2, things) {
		t.Errorf("Drop should result in %#v, got %#v", things, where2)
	}

	where3 := things.Drop(100)

	if len(where3) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where3)
	}

	where4 := Num1List{}.Drop(100)

	if len(where4) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where4)
	}
}

func TestDropNum2(t *testing.T) {
	things := Num2List{ip(1), ip(17), ip(5), ip(6), ip(17), ip(8), ip(9)}

	where1 := things.Drop(3)

	expected1 := Num2List{ip(6), ip(17), ip(8), ip(9)}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Drop should result in %#v, got %#v", expected1, where1)
	}
}

func TestDropThing(t *testing.T) {
	things := ThingList{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
		{"Boo", 4},
		{"Bam", 5},
	}

	where1 := things.Drop(3)

	expected1 := ThingList{
		{"Boo", 4},
		{"Bam", 5},
	}

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("Drop should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.Drop(10)

	if len(where2) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where2)
	}

	where3 := things.Drop(0)

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("Drop should result in %#v, got %#v", things, where3)
	}

	where4 := ThingList{}.Drop(10)

	if len(where4) != 0 {
		t.Errorf("Drop should result in empty list, got %#v", where4)
	}
}
