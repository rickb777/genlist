package main

import "testing"

func TestForeachOther(t *testing.T) {
	things := OtherSlice{60, -20, 100}

	sum := Other(0)
	things.Foreach(func(x Other) {
		sum += x
	})

	if sum != 140 {
		t.Errorf("Foreach should concatenate: %s", sum)
	}
}

func TestForeachThing(t *testing.T) {
	things := ThingSlice{
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
