package main

import "testing"

func TestContainsNum2(t *testing.T) {
	things := num2Collection(ip(50), ip(100), ip(9), ip(7), ip(100), ip(99))

	has1 := things.Contains(ip(3))

	if has1 {
		t.Errorf("Contains should not evaluate true for 3")
	}

	has2 := things.Contains(ip(7))

	if !has2 {
		t.Errorf("Contains should evaluate true for 7 in %#v", things)
	}

	has3 := num2Collection().Contains(ip(1))

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}
