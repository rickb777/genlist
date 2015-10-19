package main

import (
	"testing"
)

func TestMkStringNum1(t *testing.T) {
	things := Num1List{1, 2, 3, 5, 8, 13}

	s1 := things.MkString(", ")
	if s1 != "1, 2, 3, 5, 8, 13" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := things.MkString3("[", ", ", "]")
	if s2 != "[1, 2, 3, 5, 8, 13]" {
		t.Errorf("MkString got %q", s2)
	}
}

func TestMkStringNum2(t *testing.T) {
	things := Num2List{ip(1), ip(2), ip(3), ip(5), ip(8), ip(13)}

	s1 := things.MkString(", ")
	if s1 != "1, 2, 3, 5, 8, 13" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := things.MkString3("[", ", ", "]")
	if s2 != "[1, 2, 3, 5, 8, 13]" {
		t.Errorf("MkString got %q", s2)
	}
}
