package main

import (
	"testing"
)

func TestLastIndexWhere2Num(t *testing.T) {
	things := Num1List{1, 3, 17, 5, 6, 17, 8, 9}

	where1 := things.LastIndexWhere2(func(x Num1) bool {
		return x == 17
	}, 10)

	if where1 != 5 {
		t.Errorf("LastIndexWhere2 should be 5, got %#v", where1)
	}

	where2 := things.LastIndexWhere2(func(x Num1) bool {
		return false
	}, 4)

	if where2 != -1 {
		t.Errorf("LastIndexWhere2 should be -1, got %#v", where2)
	}

	where3 := Num1List{}.LastIndexWhere2(func(x Num1) bool {
		return true
	}, 10)

	if where3 != -1 {
		t.Errorf("LastIndexWhere2 should be -1, got %#v", where3)
	}
}

func TestLastIndexWhere2Thing(t *testing.T) {
	things := ThingList{
		{"Fee", 1},
		{"Fie", 2},
		{"Foe", 3},
		{"Boo", 4},
		{"Boo", 5},
		{"Bam", 6},
		{"Bam", 7},
	}

	where1a := things.LastIndexWhere2(func(x Thing) bool {
		return x.Name == "Boo"
	}, 10)

	if where1a != 4 {
		t.Errorf("LastIndexWhere2 should be 4, got %#v", where1a)
	}

	where1b := things.LastIndexWhere2(func(x Thing) bool {
		return x.Name == "Boo"
	}, 3)

	if where1b != 3 {
		t.Errorf("LastIndexWhere2 should be 3, got %#v", where1b)
	}

	where2 := things.LastIndexWhere2(func(x Thing) bool {
		return false
	}, 10)

	if where2 != -1 {
		t.Errorf("LastIndexWhere2 should be -1, got %#v", where2)
	}

	where3 := ThingList{}.LastIndexWhere2(func(x Thing) bool {
		return true
	}, 10)

	if where3 != -1 {
		t.Errorf("LastIndexWhere2 should be -1, got %#v", where3)
	}
}
