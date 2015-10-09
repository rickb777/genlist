package main

import "testing"

func TestCombination(t *testing.T) {
	list1 := FooList{50, 100, 9, 7, 100, 99}

	var seq1 FooSeq
	seq1 = list1

	has1 := seq1.Contains(3)
	if has1 {
		t.Errorf("Contains should not evaluate true for 3")
	}

	has2 := things.Contains(7)

	if !has2 {
		t.Errorf("Contains should evaluate true for 7")
	}

	has3 := FooList{}.Contains(1)

	if has3 {
		t.Errorf("Contains should evaluate false for empty slices")
	}
}
