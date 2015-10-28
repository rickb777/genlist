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
	things := ThingList{
		{"First", 2},
		{"Second", 3},
		{"Third", 5},
	}

	h1 := things.Last()
	t1 := things.Init().ToList()

	if h1.Number != 5 {
		t.Errorf("Last should be 5")
	}

	h2 := t1.Last()
	t2 := t1.Init().ToList()

	if h2.Number != 3 {
		t.Errorf("Last should be 3")
	}

	h3 := t2.Last()
	t3 := t2.Init()

	if h3.Number != 2 {
		t.Errorf("Last should be 2")
	}

	if t3.Size() != 0 {
		t.Errorf("Init should be an empty list")
	}
}
