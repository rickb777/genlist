package main

import "testing"

func TestForallOther(t *testing.T) {
	things := OtherList{60, -20, 100}

	all1 := things.Forall(func(x Other) bool {
		return x == 60
	})

	if all1 {
		t.Errorf("Forall should be false for 60")
	}

	all2 := things.Forall(func(x Other) bool {
		return x == 60 || x == -20 || x == 100
	})

	if !all2 {
		t.Errorf("Forall should be true")
	}

	all3 := OtherList{}.Forall(func(x Other) bool {
		return false
	})

	if !all3 {
		t.Errorf("Forall should evaluate true for empty slices")
	}
}

func TestForallThing(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

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

	all3 := ThingList{}.Forall(func(x Thing) bool {
		return false
	})

	if !all3 {
		t.Errorf("Forall should evaluate true for empty slices")
	}
}
