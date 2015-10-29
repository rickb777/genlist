package main

import "testing"

func TestSendNum1(t *testing.T) {
	a := num1Collection(1, 2, 3, 4)
	b := num1Collection()
	for v := range a.Send() {
		b = b.Add(v)
	}

	if !a.Equals(b) {
		t.Errorf("Expected the same but got %v and %v", a, b)
	}
}

func TestSendThing(t *testing.T) {
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
	tot := 0
	for v := range things.Send() {
		sum += v.Number
		tot += len(v.Name)
	}

	if sum != 53 {
		t.Errorf("Expected sum 53 but got %d", sum)
	}
	if tot != 21 {
		t.Errorf("Expected len 21 but got %d", tot)
	}
}
