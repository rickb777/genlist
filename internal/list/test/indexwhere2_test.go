package main

import (
	"testing"
)

func TestIndexWhere2Num(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9)

	where1a := things.IndexWhere2(func(x Num1) bool {
		return x == 17
	}, 2)

	if where1a != 2 {
		t.Errorf("IndexWhere2 should be 2, got %#v", where1a)
	}

	where1b := things.IndexWhere2(func(x Num1) bool {
		return x == 17
	}, 3)

	if where1b != 5 {
		t.Errorf("IndexWhere2 should be 5, got %#v", where1b)
	}

	where2 := things.IndexWhere2(func(x Num1) bool {
		return false
	}, 1)

	if where2 != -1 {
		t.Errorf("IndexWhere2 should be -1, got %#v", where2)
	}

	where3 := num1Collection().IndexWhere2(func(x Num1) bool {
		return true
	}, 0)

	if where3 != -1 {
		t.Errorf("IndexWhere2 should be -1, got %#v", where3)
	}
}

func TestIndexWhere2Thing(t *testing.T) {
	things := ThingList{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
		{"Boo", 4},
		{"Boo", 5},
		{"Bam", 6},
		{"Bam", 7},
	}

	where1a := things.IndexWhere2(func(x Thing) bool {
		return x.Name == "Boo"
	}, 1)

	if where1a != 3 {
		t.Errorf("IndexWhere2 should be 3, got %#v", where1a)
	}

	where1b := things.IndexWhere2(func(x Thing) bool {
		return x.Name == "Boo"
	}, 4)

	if where1b != 4 {
		t.Errorf("IndexWhere2 should be 4, got %#v", where1b)
	}

	where2 := things.IndexWhere2(func(x Thing) bool {
		return false
	}, 0)

	if where2 != -1 {
		t.Errorf("IndexWhere2 should be -1, got %#v", where2)
	}

	where3 := ThingList{}.IndexWhere2(func(x Thing) bool {
		return true
	}, 0)

	if where3 != -1 {
		t.Errorf("IndexWhere2 should be -1, got %#v", where3)
	}
}
