package main

import "testing"

func TestSumNum(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	number := func(x Thing) Num1 {
		return x.Number
	}

	sum1 := things.SumNum1(number)

	if sum1 != 140 {
		t.Errorf("SumNum should result in %v, got %v", 340, sum1)
	}
}
