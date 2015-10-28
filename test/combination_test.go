package main

import "testing"

func TestCombination(t *testing.T) {
	list1 := Foo1List{50, 100, 9, 7, 100, 99}

	var seq1 Foo1Collection
	seq1 = list1

	if seq1.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	if !seq1.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

//	has2 := things.Contains(7)
//
//	if !has2 {
//		t.Errorf("Contains should evaluate true for 7")
//	}
//
//	has3 := FooList{}.Contains(1)
//
//	if has3 {
//		t.Errorf("Contains should evaluate false for empty slices")
//	}
}
