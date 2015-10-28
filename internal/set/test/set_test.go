package main

import "testing"

func TestLenFoo(t *testing.T) {
	emptyThing := fooCollection()
	if emptyThing.Size() != 0 {
		t.Errorf("Size should be 0")
	}
	if !emptyThing.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}
	if emptyThing.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}

	someThing := fooCollection("aa", "bb")
	if someThing.Size() != 2 {
		t.Errorf("Size should be 2")
	}
	if someThing.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}
	if !someThing.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}
}

func TestLenNum1(t *testing.T) {
	emptyThing := num1Collection()
	if emptyThing.Size() != 0 {
		t.Errorf("Size should be 0")
	}
	if !emptyThing.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}
	if emptyThing.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}

	someThing := num1Collection(1, 2)
	if someThing.Size() != 2 {
		t.Errorf("Size should be 2")
	}
	if someThing.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}
	if !someThing.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}
}

func TestToSlice(t *testing.T) {
	a := num1Collection(1, 2, 3)
	b := a.ToSlice()

	if len(b) != 3 {
		t.Errorf("Expected length 3 in %v", b)
	}

	if b[0] != 1 && b[0] != 2 && b[0] != 3 {
		t.Errorf("Expected first element 1, 2 or 3 in %v", b)
	}
}

func TestAddRemoveContains(t *testing.T) {
	a := num1Collection()

	b := a.Add(71)

	if !b.Contains(71) {
		t.Errorf("Contains should contain 71")
	}

	c := b.Remove(71)

	if c.Contains(71) {
		t.Errorf("Contains should not contain 71")
	}

	d := c.Add(13, 7, 1)

	if !(d.Contains(13) && d.Contains(7) && d.Contains(1)) {
		t.Errorf("Contains should contain 13, 7, 1")
	}
}

func TestSetHasNoDuplicate(t *testing.T) {
	a := num1Collection(3, 5, 8, 3)

	if a.Size() != 3 {
		t.Errorf("set should have 3 elements with no duplicates")
	}

	if !(a.Contains(8) && a.Contains(5) && a.Contains(3)) {
		t.Errorf("set should contain 3, 5, 8: %v", a)
	}
}

func TestContainsAll(t *testing.T) {
	a := num1Collection(2, 3, 5, 8, 13, 21)

	if !a.ContainsAll(2, 3, 5, 8, 13, 21) {
		t.Errorf("ContainsAll should contain 6 fibs")
	}

	if a.ContainsAll(2, 3, 5, 8, 13, 20) {
		t.Errorf("ContainsAll should not contain non-fibs")
	}
}

func TestSetIsSubsetAndSuperset(t *testing.T) {
	a := num1Collection(1, 2, 3, 5, 7)
	b := num1Collection(1, 3, 7)
	c := num1Collection(1, 3, 5, 7, 9)
	e := num1Collection()

	if !a.Equals(a) {
		t.Errorf("set %v should be equal to itself", a)
	}
	if !a.IsSubset(a) {
		t.Errorf("set %v should be a subset of itself", a)
	}

	if a.IsProperSubset(a) {
		t.Errorf("set %v should not be a proper subset of itself", a)
	}

	if b.Equals(a) {
		t.Errorf("set %v should not equal %v", b, a)
	}
	if b.Equals(c) {
		t.Errorf("set %v should not equal %v", b, c)
	}
	if !b.IsSubset(c) {
		t.Errorf("set %v should be a subset of set %v", b, c)
	}
	if !b.IsSubset(a) {
		t.Errorf("set %v should be a subset of set %v", b, a)
	}
	if !b.IsProperSubset(a) {
		t.Errorf("set %v should be a proper subset of set %v", b, a)
	}
	if !a.IsSuperset(b) {
		t.Errorf("set %v should be a superset of set %v", a, b)
	}

	if !e.IsSubset(a) {
		t.Errorf("set %v should be a subset of set %v", e, a)
	}
	if !e.IsProperSubset(a) {
		t.Errorf("set %v should be a proper subset of set %v", e, a)
	}
	if !a.IsSuperset(e) {
		t.Errorf("set %v should be a superset of set %v", e, a)
	}

	d := num1Collection(1, 3, 4, 7)

	if d.IsSubset(a) {
		t.Errorf("set %v should not be a subset of set %v", d, a)
	}
	if d.IsProperSubset(a) {
		t.Errorf("set %v should not be a proper subset of set %v", d, a)
	}
	if a.IsSuperset(d) {
		t.Errorf("set %v should not be a superset of set %v", a, d)
	}
}

func TestUnion(t *testing.T) {
	a := num1Collection(1, 2, 3, 5, 8)
	b := num1Collection(4, 5, 6, 7, 8)
	e := num1Collection()

	u1 := a.Union(e)
	if u1.Size() != 5 {
		t.Errorf("expected size 5, got %v", u1)
	}

	u2 := e.Union(a)
	if u1.Size() != 5 {
		t.Errorf("expected size 5, got %v", u2)
	}

	u3 := b.Union(a)
	if u3.Size() != 8 {
		t.Errorf("expected size 8, got %v", u3)
	}
}

func TestIntersection(t *testing.T) {
	a := num1Collection(1, 2, 3, 5, 8)
	b := num1Collection(4, 5, 6, 7, 8)
	c := num1Collection(4, 9, 16, 25)
	e := num1Collection()

	i1 := a.Intersection(e)
	if i1.Size() != 0 {
		t.Errorf("expected empty set, got %v", i1)
	}

	i2 := a.Intersection(b)
	if i2.Size() != 2 {
		t.Errorf("expected 2 common items, got %v", i2)
	}

	i3 := a.Intersection(a)
	if i3.Size() != 5 {
		t.Errorf("expected 5 items, got %v", i3)
	}

	i4 := c.Intersection(b)
	if i4.Size() != 1 {
		t.Errorf("expected 1 item, got %v", i4)
	}
}

func TestDifference(t *testing.T) {
	a := num1Collection(1, 2, 3)
	b := num1Collection(1, 3, 4, 5, 6, 10)

	d1 := a.Difference(b)

	if !(d1.Size() == 1 && d1.Contains(2)) {
		t.Errorf("the difference is wrong: %v", d1)
	}

	d2 := b.Difference(a)

	if !(d2.Size() == 4 && d2.Equals(num1Collection(4, 5, 6, 10))) {
		t.Errorf("the difference is wrong: %v", d2)
	}
}

func TestSend(t *testing.T) {
	a := num1Collection(1, 2, 3, 4)
	b := num1Collection()
	for v := range a.Send() {
		b = b.Add(v)
	}

	if !a.Equals(b) {
		t.Errorf("Expected the same but got %v and %v", a, b)
	}
}

func TestForall(t *testing.T) {
	someThing := num1Collection(10)
	if !someThing.Forall(func(x Num1) bool { return x > 0 }) {
		t.Errorf("Forall should be true")
	}

	if someThing.Forall(func(x Num1) bool { return x < 0 }) {
		t.Errorf("Forall should be false")
	}
}

func TestFilter(t *testing.T) {
	someThing := num1Collection(10)
	v1 := someThing.Filter(func(x Num1) bool { return true })
	if !v1.Equals(someThing) {
		t.Errorf("Filter failed: got %v", v1)
	}

	v2 := someThing.Filter(func(x Num1) bool { return false })
	if v2.NonEmpty() {
		t.Errorf("Filter should be empty")
	}
}

func TestPartition(t *testing.T) {
	someThing := num1Collection(10)
	m1, o1 := someThing.Partition(func(Num1) bool { return true })
	if !m1.Equals(someThing) {
		t.Errorf("Partition match should be 10")
	}

	if o1.NonEmpty() {
		t.Errorf("Partition other should be empty")
	}

	m2, o2 := someThing.Partition(func(Num1) bool { return false })
	if m2.NonEmpty() {
		t.Errorf("Partition match should be empty")
	}

	if !o2.Equals(someThing) {
		t.Errorf("Partition other should be 10")
	}
}

func TestForeach(t *testing.T) {
	someThing := num1Collection(1, 2, 3)
	x := Num1(0)
	someThing.Foreach(func(v Num1) { x += v })
	if x != 6 {
		t.Errorf("Foreach should set x to 6 not %v", x)
	}
}
