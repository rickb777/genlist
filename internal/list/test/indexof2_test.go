package main

import (
	"testing"
)

func TestIndexOf2Num1(t *testing.T) {
	things := Num1List{1, 3, 17, 5, 6, 17, 8, 9}

	where1a := things.IndexOf2(17, 2)

	if where1a != 2 {
		t.Errorf("IndexOf should be 2, got %#v", where1a)
	}

	where1b := things.IndexOf2(17, 3)

	if where1b != 5 {
		t.Errorf("IndexOf should be 5, got %#v", where1b)
	}

	where2 := things.IndexOf2(19, 0)

	if where2 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where2)
	}

	where3 := Num1List{}.IndexOf2(1, 0)

	if where3 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where3)
	}
}

func TestIndexOf2Num2(t *testing.T) {
	things := Num2List{ip(1), ip(3), ip(17), ip(5), ip(6), ip(17), ip(8), ip(9)}

	where1a := things.IndexOf2(ip(17), 2)

	if where1a != 2 {
		t.Errorf("IndexOf should be 2, got %#v", where1a)
	}

	where1b := things.IndexOf2(ip(17), 3)

	if where1b != 5 {
		t.Errorf("IndexOf should be 5, got %#v", where1b)
	}

	where2 := things.IndexOf2(ip(19), 3)

	if where2 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where2)
	}

	where3 := Num2List{}.IndexOf2(ip(1), 0)

	if where3 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where3)
	}
}
