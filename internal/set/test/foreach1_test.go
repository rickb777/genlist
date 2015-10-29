package main

import (
	"testing"
)

func TestForeachNum1(t *testing.T) {
	things := num1Collection(60, -20, 100)

	sum := Num1(0)
	things.Foreach(func(x Num1) {
		sum += x
	})

	if sum != 140 {
		t.Errorf("Foreach should concatenate: %s", sum)
	}
}

func TestForeachThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	sum := Num1(0)
	things.Foreach(func(x Thing) {
		sum += x.Number
	})

	if sum != 53 {
		t.Errorf("Foreach should concatenate: %d", 53)
	}
}
