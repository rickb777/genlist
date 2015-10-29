package main

import "testing"

func TestShuffleNum(t *testing.T) {
	original := num1Collection(50, 100, 9, 7, 100, 99)

	shuffled := original.Shuffle()

	notFound := func(t Num1) bool {
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
	original := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

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
