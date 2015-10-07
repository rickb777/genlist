package main

import "testing"

func TestForeach(t *testing.T) {
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
