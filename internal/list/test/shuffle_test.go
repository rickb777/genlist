package main

import "testing"

func TestShuffleOther(t *testing.T) {
	original := OtherList{50, 100, 9, 7, 100, 99}

	shuffled := original.Shuffle()

	notFound := func(t Other) bool {
		for x := 0; x < len(shuffled); x++ {
			if shuffled[x] == t {
				return false
			}
		}
		return true
	}

	if original.Exists(notFound) {
		t.Error("The shuffled list is missing elements")
	}
}

func TestShuffleThing(t *testing.T) {
	original := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	shuffled := original.Shuffle()

	notFound := func(t Thing) bool {
		for x := 0; x < len(shuffled); x++ {
			if shuffled[x] == t {
				return false
			}
		}
		return true
	}

	if original.Exists(notFound) {
		t.Error("The shuffled list is missing elements")
	}
}
