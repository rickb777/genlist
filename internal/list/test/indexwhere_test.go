package main

import (
	"testing"
)

func TestIndexWhereNum(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9)

	where1 := things.IndexWhere(func(x Num1) bool {
		return x == 17
	})

	if where1 != 2 {
		t.Errorf("IndexWhere should be 2, got %#v", where1)
	}

	where2 := things.IndexWhere(func(x Num1) bool {
		return false
	})

	if where2 != -1 {
		t.Errorf("IndexWhere should be -1, got %#v", where2)
	}

	where3 := num1Collection().IndexWhere(func(x Num1) bool {
		return true
	})

	if where3 != -1 {
		t.Errorf("IndexWhere should be -1, got %#v", where3)
	}
}

func TestIndexWhereThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.IndexWhere(func(x Thing) bool {
		return x.Name == "Boo"
	})

	if where1 != 3 {
		t.Errorf("IndexWhere should be 3, got %#v", where1)
	}

	where2 := things.IndexWhere(func(x Thing) bool {
		return false
	})

	if where2 != -1 {
		t.Errorf("IndexWhere should be -1, got %#v", where2)
	}

	where3 := ThingList{}.IndexWhere(func(x Thing) bool {
		return true
	})

	if where3 != -1 {
		t.Errorf("IndexWhere should be -1, got %#v", where3)
	}
}
