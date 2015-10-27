package main

import "testing"

func TestExistsNum2(t *testing.T) {
	things := num2Collection(ip(60), ip(-20), ip(100))

	any1 := things.Exists(func(x *Num2) bool {
		return *x == 10
	})

	if any1 {
		t.Errorf("Exists should not evaluate true for 10")
	}

	any2 := things.Exists(func(x *Num2) bool {
		return *x > 50
	})

	if !any2 {
		t.Errorf("Exists should evaluate true for Number > 50")
	}

	any3 := num2Collection().Exists(func(x *Num2) bool {
		return true
	})

	if any3 {
		t.Errorf("Exists should evaluate false for empty slices")
	}
}
