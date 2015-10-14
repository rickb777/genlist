package main

import "testing"

func TestForeachNum1(t *testing.T) {
	things := Num1List{60, -20, 100}

	sum := Num1(0)
	things.Foreach(func(x Num1) {
		sum += x
	})

	if sum != 140 {
		t.Errorf("Foreach should concatenate: %s", sum)
	}
}

func TestForeachThing(t *testing.T) {
	things := ThingList{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	concat := ""
	things.Foreach(func(x Thing) {
		concat += x.Name
	})

	if concat != "FirstSecondThird" {
		t.Errorf("Foreach should concatenate: %s", concat)
	}
}
