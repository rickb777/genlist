package main

import "testing"

func TestForallNum1(t *testing.T) {
	things := num1Collection(60, -20, 100)

	all1 := things.Forall(func(x Num1) bool {
		return x == 60
	})

	if all1 {
		t.Errorf("Forall should be false for 60")
	}

	all2 := things.Forall(func(x Num1) bool {
		return x == 60 || x == -20 || x == 100
	})

	if !all2 {
		t.Errorf("Forall should be true")
	}

	all3 := num1Collection().Forall(func(x Num1) bool {
		return false
	})

	if !all3 {
		t.Errorf("Forall should evaluate true for empty slices")
	}
}

func TestForallThing(t *testing.T) {
	things := thingCollection(
		Thing{"First", 60},
		Thing{"Second", -20},
		Thing{"Third", 100},
	)

	all1 := things.Forall(func(x Thing) bool {
		return x.Name == "First"
	})

	if all1 {
		t.Errorf("Forall should be false for Name == 'First'")
	}

	all2 := things.Forall(func(x Thing) bool {
		return x.Name == "First" || x.Name == "Second" || x.Name == "Third"
	})

	if !all2 {
		t.Errorf("Forall should be true")
	}

	all3 := thingCollection().Forall(func(x Thing) bool {
		return false
	})

	if !all3 {
		t.Errorf("Forall should evaluate true for empty slices")
	}
}
