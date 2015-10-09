package main

import "testing"

func TestExistsOther(t *testing.T) {
	things := OtherList{60, -20, 100}

	any1 := things.Exists(func(x Other) bool {
		return x == 10
	})

	if any1 {
		t.Errorf("Exists should not evaluate true for 10")
	}

	any2 := things.Exists(func(x Other) bool {
		return x > 50
	})

	if !any2 {
		t.Errorf("Exists should evaluate true for Number > 50")
	}

	any3 := OtherList{}.Exists(func(x Other) bool {
		return true
	})

	if any3 {
		t.Errorf("Exists should evaluate false for empty slices")
	}
}

func TestExistsThing(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	any1 := things.Exists(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if any1 {
		t.Errorf("Exists should not evaluate true for Name == Dummy")
	}

	any2 := things.Exists(func(x Thing) bool {
		return x.Number > 50
	})

	if !any2 {
		t.Errorf("Exists should evaluate true for Number > 50")
	}

	any3 := ThingList{}.Exists(func(x Thing) bool {
		return true
	})

	if any3 {
		t.Errorf("Exists should evaluate false for empty slices")
	}
}
