package main

import (
	"testing"
)

func TestMkStringNum1(t *testing.T) {
	things := num1Collection(1, 2, 3, 5, 8, 13)

	s1 := things.MkString3("[", ", ", "]")
	e1 := "[1, 2, 3, 5, 8, 13]"
	if s1 != e1 {
		t.Errorf("MkString got %q", s1)
	}

	s2 := things.MkString(", ")
	e2 := "1, 2, 3, 5, 8, 13"
	if s2 != e2 {
		t.Errorf("MkString got %q", s2)
	}

	s3 := things.String()
	e3 := "[1,2,3,5,8,13]"
	if s3 != e3 {
		t.Errorf("MkString got %q", s3)
	}
}

func TestMkStringNum2(t *testing.T) {
	things := num2Collection(ip(1), ip(2), ip(3), ip(5), ip(8), ip(13))

	s1 := things.MkString(", ")
	if s1 != "1, 2, 3, 5, 8, 13" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := things.MkString3("[", ", ", "]")
	if s2 != "[1, 2, 3, 5, 8, 13]" {
		t.Errorf("MkString got %q", s2)
	}

	s3 := things.String()
	if s3 != "[1,2,3,5,8,13]" {
		t.Errorf("MkString got %q", s3)
	}
}

func TestMkStringThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	s1 := things.MkString3("[", ", ", "]")
	e1 := "[{Fee 1}, {Fie 2}, {Foe 3}, {Boo 5}, {Boo 8}, {Bam 13}, {Bam 21}]"
	if s1 != e1 {
		t.Errorf("MkString got %q not %q", s1, e1)
	}

	s2 := things.MkString(", ")
	e2 := "{Fee 1}, {Fie 2}, {Foe 3}, {Boo 5}, {Boo 8}, {Bam 13}, {Bam 21}"
	if s2 != e2 {
		t.Errorf("MkString got %q not %q", s2, e2)
	}

	s3 := things.String()
	e3 := "[{Fee 1},{Fie 2},{Foe 3},{Boo 5},{Boo 8},{Bam 13},{Bam 21}]"
	if s3 != e3 {
		t.Errorf("MkString got %q not %q", s3, e3)
	}
}

