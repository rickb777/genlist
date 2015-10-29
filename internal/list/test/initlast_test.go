package main

import "testing"

func TestInitLastNum1(t *testing.T) {
	things := num1Collection(3, 5, 8, 13)

	h1 := things.Last()
	t1 := things.Init().ToList()

	if h1 != 13 {
		t.Errorf("Last should be 13")
	}

	h2 := t1.Last()
	t2 := t1.Init().ToList()

	if h2 != 8 {
		t.Errorf("Last should be 8")
	}

	h3 := t2.Last()
	t3 := t2.Init().ToList()

	if h3 != 5 {
		t.Errorf("Last should be 5")
	}

	h4 := t3.Last()
	t4 := t3.Init()

	if h4 != 3 {
		t.Errorf("Last should be 3")
	}

	if t4.Size() != 0 {
		t.Errorf("Init should be an empty list")
	}
}

func TestInitLastNum2(t *testing.T) {
	things := num2Collection(ip(3), ip(5), ip(8), ip(13))

	h1 := things.Last()
	t1 := things.Init().ToList()

	if *h1 != 13 {
		t.Errorf("Last should be 13")
	}

	h2 := t1.Last()
	t2 := t1.Init().ToList()

	if *h2 != 8 {
		t.Errorf("Last should be 8")
	}

	h3 := t2.Last()
	t3 := t2.Init().ToList()

	if *h3 != 5 {
		t.Errorf("Last should be 5")
	}

	h4 := t3.Last()
	t4 := t3.Init()

	if *h4 != 3 {
		t.Errorf("Last should be 3")
	}

	if t4.Size() != 0 {
		t.Errorf("Init should be an empty list")
	}
}

func TestInitLastThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	h1 := things.Last()
	t1 := things.Init().ToList()

	if h1.Number != 21 {
		t.Errorf("Last should be 21")
	}

	h2 := t1.Last()
	t2 := t1.Init().ToList()

	if h2.Number != 13 {
		t.Errorf("Last should be 13")
	}

	h3 := t2.Last()
	t3 := t2.Init()

	if h3.Number != 8 {
		t.Errorf("Last should be 8")
	}

	if t3.Size() != 4 {
		t.Errorf("Init should leave 4 items not %d", t3.Size())
	}
}
