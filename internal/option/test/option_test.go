package main

import (
	"testing"
	"fmt"
)

func TestLen(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Size() != 1 {
		t.Errorf("Size should be 1")
	}

	noThing := NoOther()
	if noThing.Size() != 0 {
		t.Errorf("Size should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	noThing := NoOther()
	if !noThing.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}
}

func TestNonEmpty(t *testing.T) {
	someThing := SomeOther(60)
	if !someThing.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

	noThing := NoOther()
	if noThing.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}
}

func TestGet(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Get() != 60 {
		t.Errorf("Get should be 60")
	}
}

func TestGetOrElse(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.GetOrElse(func() Other { return 50 }) != 60 {
		t.Errorf("GetOrElse should be 60")
	}

	noThing := NoOther()
	if noThing.GetOrElse(func() Other { return 50 }) != 50 {
		t.Errorf("GetOrElse should be 50")
	}
}

func TestOrElse(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.OrElse(func() OptionalOther { return NoOther() }).Get() != 60 {
		t.Errorf("OrElse should be 60")
	}

	noThing := NoOther()
	if noThing.OrElse(func() OptionalOther { return SomeOther(60) }).Get() != 60 {
		t.Errorf("OrElse should be 60")
	}
}

func TestExists(t *testing.T) {
	someThing := SomeOther(60)
	if !someThing.Exists(func(Other) bool { return true }) {
		t.Errorf("Exists should be true")
	}

	noThing := NoOther()
	if noThing.Exists(func(Other) bool { return true }) {
		t.Errorf("Exists should be false")
	}
}

func TestForall(t *testing.T) {
	someThing := SomeOther(60)
	if !someThing.Forall(func(Other) bool { return true }) {
		t.Errorf("Forall should be true")
	}

	noThing := NoOther()
	if !noThing.Forall(func(Other) bool { return true }) {
		t.Errorf("Forall should be true")
	}
}

func TestFilter(t *testing.T) {
	someThing := SomeOther(60)
	v1 := someThing.Filter(func(Other) bool { return true })
	if v1.(OptionalOther).Get() != 60 {
		t.Errorf("Filter should be 60")
	}

	v2 := someThing.Filter(func(Other) bool { return false })
	if v2.NonEmpty() {
		t.Errorf("Filter should be empty")
	}

	noThing := NoOther()
	v3 := noThing.Filter(func(Other) bool { return true })
	if v3.NonEmpty() {
		t.Errorf("Filter should be empty")
	}
}

func TestPartition(t *testing.T) {
	someThing := SomeOther(60)
	m1, o1 := someThing.Partition(func(Other) bool { return true })
	if m1.(OptionalOther).Get() != 60 {
		t.Errorf("Partition match should be 60")
	}

	if o1.NonEmpty() {
		t.Errorf("Partition other should be empty")
	}

	m2, o2 := someThing.Partition(func(Other) bool { return false })
	if m2.NonEmpty() {
		t.Errorf("Partition match should be empty")
	}

	if o2.(OptionalOther).Get() != 60 {
		t.Errorf("Partition other should be 60")
	}

	noThing := NoOther()
	m3, o3 := noThing.Partition(func(Other) bool { return true })
	if m3.NonEmpty() {
		t.Errorf("Partition match should be empty")
	}
	if o3.NonEmpty() {
		t.Errorf("Partition other should be empty")
	}
}

func TestFind(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Find(func(Other) bool { return true }).Get() != 60 {
		t.Errorf("Find should be 60")
	}

	noThing := NoOther()
	if noThing.Find(func(Other) bool { return true }).NonEmpty() {
		t.Errorf("Find should be empty")
	}
}

func TestForeach(t *testing.T) {
	someThing := SomeOther(60)
	x := Other(0)
	someThing.Foreach(func(Other) { x = Other(1) })
	if x != 1 {
		t.Errorf("Foreach should set x")
	}

	noThing := NoOther()
	y := Other(0)
	noThing.Foreach(func(Other) { y = Other(1) })
	if y != 0 {
		t.Errorf("Foreach should not set x")
	}
}

func TestContains(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Contains(50) {
		t.Errorf("Should not contain 50")
	}

	if !someThing.Contains(60) {
		t.Errorf("Should contain 60")
	}

	noThing := NoOther()
	if noThing.Contains(50) {
		t.Errorf("Should not contain 50")
	}
}

func TestEquals(t *testing.T) {
	someThing1 := SomeOther(10)
	someThing2 := SomeOther(20)
	noThing := NoOther()

	if someThing1.Equals(someThing2) {
		t.Errorf("Should not be equal")
	}

	if !someThing1.Equals(someThing1) {
		t.Errorf("Should be equal")
	}

	if someThing1.Equals(noThing) {
		t.Errorf("Should not be equal")
	}

	if !noThing.Equals(noThing) {
		t.Errorf("Should not be equal")
	}

	if noThing.Equals(someThing1) {
		t.Errorf("Should not be equal")
	}
}

func TestCount(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Count(50) != 0 {
		t.Errorf("Should contain zero 50s")
	}

	if someThing.Count(60) != 1 {
		t.Errorf("Should contain one 60")
	}

	noThing := NoOther()
	if noThing.Count(50) != 0 {
		t.Errorf("Should not contain 50")
	}
}

func TestDistinct(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Distinct().Size() != 1 {
		t.Errorf("Should be length 1")
	}

	noThing := NoOther()
	if noThing.Distinct().NonEmpty() {
		t.Errorf("Should be empty")
	}
}

func TestSum(t *testing.T) {
	someThing := SomeOther(60)
	if someThing.Sum() != 60 {
		t.Errorf("Sum should be 60")
	}

	noThing := NoOther()
	if noThing.Sum() != 0 {
		t.Errorf("Sum should be 0")
	}
}

func TestMapTo(t *testing.T) {
	someThing := SomeOther(60)
	fn1 := func(o Other) Foo { return Foo(fmt.Sprintf("%d", o)) }
	m1 := someThing.MapToFoo(fn1)
	if m1.Head() != "60" {
		t.Errorf("MapToFoo should be '60' but got %q", m1)
	}

	noThing := NoOther()
	m2 := noThing.MapToFoo(fn1)
	if m2.NonEmpty() {
		t.Errorf("MapToFoo should be absent but got %+v", m2)
	}
}

func TestFlatMapTo(t *testing.T) {
	someThing := SomeOther(60)
	fn0 := func(o Other) FooSeq { return NoFoo() }
	fn1 := func(o Other) FooSeq { return SomeFoo(Foo(fmt.Sprintf("%d", o))) }

	m0 := someThing.FlatMapToFoo(fn0)
	if m0.NonEmpty() {
		t.Errorf("MapToFoo should be absent but got %q", m0)
	}

	m1 := someThing.FlatMapToFoo(fn1)
	if m1.Head() != "60" {
		t.Errorf("MapToFoo should be '60' but got %q", m1)
	}

	noThing := NoOther()
	m2 := noThing.FlatMapToFoo(fn1)
	if m2.NonEmpty() {
		t.Errorf("MapToFoo should be absent but got %+v", m2)
	}
}

func TestMkString1(t *testing.T) {
	someThing := SomeOther(123)

	s1 := someThing.MkString3("[", ", ", "]")
	if s1 != "[123]" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := someThing.MkString(", ")
	if s2 != "[123]" {
		t.Errorf("MkString got %q", s2)
	}

	s3 := someThing.String()
	if s3 != "[123]" {
		t.Errorf("MkString got %q", s3)
	}
}
