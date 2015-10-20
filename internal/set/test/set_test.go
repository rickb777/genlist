package main

import "testing"

func TestLen(t *testing.T) {
	someThing := NewFooSet("aa")
	if someThing.Size() != 1 {
		t.Errorf("Size should be 1")
	}

}
