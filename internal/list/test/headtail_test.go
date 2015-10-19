package main

import "testing"

func TestHeadTailNum1(t *testing.T) {
	things := Num1List{3, 5, 8, 13}

	h1 := things.Head()
	t1 := things.Tail()

	if h1 != 3 {
		t.Errorf("Head should be 3")
	}

	h2 := t1.Head()
	t2 := t1.Tail()

	if h2 != 5 {
		t.Errorf("Head should be 5")
	}

	h3 := t2.Head()
	t3 := t2.Tail()

	if h3 != 8 {
		t.Errorf("Head should be 8")
	}

	h4 := t3.Head()
	t4 := t3.Tail()

	if h4 != 13 {
		t.Errorf("Head should be 13")
	}

	if t4.Len() != 0 {
		t.Errorf("Tail should be an empty list")
	}
}

func TestHeadTailNum2(t *testing.T) {
	things := Num2List{ip(3), ip(5), ip(8), ip(13)}

	h1 := things.Head()
	t1 := things.Tail()

	if *h1 != 3 {
		t.Errorf("Head should be 3")
	}

	h2 := t1.Head()
	t2 := t1.Tail()

	if *h2 != 5 {
		t.Errorf("Head should be 5")
	}

	h3 := t2.Head()
	t3 := t2.Tail()

	if *h3 != 8 {
		t.Errorf("Head should be 8")
	}

	h4 := t3.Head()
	t4 := t3.Tail()

	if *h4 != 13 {
		t.Errorf("Head should be 13")
	}

	if t4.Len() != 0 {
		t.Errorf("Tail should be an empty list")
	}
}

func TestHeadTailThing(t *testing.T) {
	things := ThingList{
		{"First", 2},
		{"Second", 3},
		{"Third", 5},
	}

	h1 := things.Head()
	t1 := things.Tail()

	if h1.Number != 2 {
		t.Errorf("Head should be 3")
	}

	h2 := t1.Head()
	t2 := t1.Tail()

	if h2.Number != 3 {
		t.Errorf("Head should be 5")
	}

	h3 := t2.Head()
	t3 := t2.Tail()

	if h3.Number != 5 {
		t.Errorf("Head should be 8")
	}

	if t3.Len() != 0 {
		t.Errorf("Tail should be an empty list")
	}
}
