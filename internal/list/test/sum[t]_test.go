package main

import "testing"

func TestSumNum(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	number := func(x Thing) Num1 {
		return x.Number
	}

	sum1 := things.SumNum1(number)

	if sum1 != 53 {
		t.Errorf("SumNum should result in 53, got %d", sum1)
	}
}
