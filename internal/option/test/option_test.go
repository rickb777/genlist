package main

import (
	"testing"
	"fmt"
)

func TestSomeNum1SimpleCases(t *testing.T) {
	someNum1 := SomeNum1(60)
	if someNum1.Size() != 1 {
		t.Errorf("Size should be 1")
	}
	if someNum1.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}
	if !someNum1.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}
	if someNum1.Get() != 60 {
		t.Errorf("Get should be 60")
	}
	if someNum1.GetOrElse(func() Num1 { return 50 }) != 60 {
		t.Errorf("GetOrElse should be 60")
	}
	if someNum1.OrElse(func() OptionalNum1 { return NoNum1() }).Get() != 60 {
		t.Errorf("OrElse should be 60")
	}
	if !someNum1.Exists(func(Num1) bool { return true }) {
		t.Errorf("Exists should be true")
	}
	if !someNum1.Forall(func(Num1) bool { return true }) {
		t.Errorf("Forall should be true")
	}
	if someNum1.Find(func(Num1) bool { return true }).Get() != 60 {
		t.Errorf("Find should be 60")
	}
	if someNum1.Sum() != 60 {
		t.Errorf("Sum should be 60")
	}
}

func TestNoNum1SimpleCases(t *testing.T) {
	noNum1 := NoNum1()
	if noNum1.Size() != 0 {
		t.Errorf("Size should be 0")
	}
	if !noNum1.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}
	if noNum1.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}
	if noNum1.GetOrElse(func() Num1 { return 50 }) != 50 {
		t.Errorf("GetOrElse should be 50")
	}
	if noNum1.OrElse(func() OptionalNum1 { return SomeNum1(60) }).Get() != 60 {
		t.Errorf("OrElse should be 60")
	}
	if noNum1.Exists(func(Num1) bool { return true }) {
		t.Errorf("Exists should be false")
	}
	if !noNum1.Forall(func(Num1) bool { return true }) {
		t.Errorf("Forall should be true")
	}
	if noNum1.Find(func(Num1) bool { return true }).NonEmpty() {
		t.Errorf("Find should be empty")
	}
	if noNum1.Sum() != 0 {
		t.Errorf("Sum should be 0")
	}
}

func TestSomeFooSimpleCases(t *testing.T) {
	someFoo := SomeFoo("foo")
	if someFoo.Size() != 1 {
		t.Errorf("Size should be 1")
	}
	if someFoo.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}
	if !someFoo.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}
	if someFoo.Get() != "foo" {
		t.Errorf("Get should be foo")
	}
	if someFoo.GetOrElse(func() Foo { return "bar" }) != "foo" {
		t.Errorf("GetOrElse should be foo")
	}
	if someFoo.OrElse(func() OptionalFoo { return NoFoo() }).Get() != "foo" {
		t.Errorf("OrElse should be foo")
	}
	if !someFoo.Exists(func(Foo) bool { return true }) {
		t.Errorf("Exists should be true")
	}
	if !someFoo.Forall(func(Foo) bool { return true }) {
		t.Errorf("Forall should be true")
	}
	if someFoo.Find(func(Foo) bool { return true }).Get() != "foo" {
		t.Errorf("Find should be foo")
	}
}

func TestNoFooSimpleCases(t *testing.T) {
	noFoo := NoFoo()
	if noFoo.Size() != 0 {
		t.Errorf("Size should be 0")
	}
	if !noFoo.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}
	if noFoo.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}
	if noFoo.GetOrElse(func() Foo { return "bar" }) != "bar" {
		t.Errorf("GetOrElse should be bar")
	}
	if noFoo.OrElse(func() OptionalFoo { return SomeFoo("bar") }).Get() != "bar" {
		t.Errorf("OrElse should be bar")
	}
	if noFoo.Exists(func(Foo) bool { return true }) {
		t.Errorf("Exists should be false")
	}
	if !noFoo.Forall(func(Foo) bool { return true }) {
		t.Errorf("Forall should be true")
	}
	if noFoo.Find(func(Foo) bool { return true }).NonEmpty() {
		t.Errorf("Find should be empty")
	}
}

func TestFilter(t *testing.T) {
	someNum1 := SomeNum1(60)
	v1 := someNum1.Filter(func(Num1) bool { return true })
	if v1.(OptionalNum1).Get() != 60 {
		t.Errorf("Filter should be 60")
	}

	v2 := someNum1.Filter(func(Num1) bool { return false })
	if v2.NonEmpty() {
		t.Errorf("Filter should be empty")
	}

	noNum1 := NoNum1()
	v3 := noNum1.Filter(func(Num1) bool { return true })
	if v3.NonEmpty() {
		t.Errorf("Filter should be empty")
	}
}

func TestPartition(t *testing.T) {
	someNum1 := SomeNum1(60)
	m1, o1 := someNum1.Partition(func(Num1) bool { return true })
	if m1.(OptionalNum1).Get() != 60 {
		t.Errorf("Partition match should be 60")
	}

	if o1.NonEmpty() {
		t.Errorf("Partition other should be empty")
	}

	m2, o2 := someNum1.Partition(func(Num1) bool { return false })
	if m2.NonEmpty() {
		t.Errorf("Partition match should be empty")
	}

	if o2.(OptionalNum1).Get() != 60 {
		t.Errorf("Partition other should be 60")
	}

	noNum1 := NoNum1()
	m3, o3 := noNum1.Partition(func(Num1) bool { return true })
	if m3.NonEmpty() {
		t.Errorf("Partition match should be empty")
	}
	if o3.NonEmpty() {
		t.Errorf("Partition other should be empty")
	}
}

func TestForeach(t *testing.T) {
	someNum1 := SomeNum1(60)
	x := Num1(0)
	someNum1.Foreach(func(Num1) { x = Num1(1) })
	if x != 1 {
		t.Errorf("Foreach should set x")
	}

	noNum1 := NoNum1()
	y := Num1(0)
	noNum1.Foreach(func(Num1) { y = Num1(1) })
	if y != 0 {
		t.Errorf("Foreach should not set x")
	}
}

func TestContains(t *testing.T) {
	someNum1 := SomeNum1(60)
	if someNum1.Contains(50) {
		t.Errorf("Should not contain 50")
	}

	if !someNum1.Contains(60) {
		t.Errorf("Should contain 60")
	}

	noNum1 := NoNum1()
	if noNum1.Contains(50) {
		t.Errorf("Should not contain 50")
	}
}

func TestEquals(t *testing.T) {
	someThing1 := SomeNum1(10)
	someThing2 := SomeNum1(20)
	noNum1 := NoNum1()

	if someThing1.Equals(someThing2) {
		t.Errorf("Should not be equal")
	}

	if !someThing1.Equals(someThing1) {
		t.Errorf("Should be equal")
	}

	if someThing1.Equals(noNum1) {
		t.Errorf("Should not be equal")
	}

	if !noNum1.Equals(noNum1) {
		t.Errorf("Should not be equal")
	}

	if noNum1.Equals(someThing1) {
		t.Errorf("Should not be equal")
	}
}

func TestCount(t *testing.T) {
	someNum1 := SomeNum1(60)
	if someNum1.Count(50) != 0 {
		t.Errorf("Should contain zero 50s")
	}

	if someNum1.Count(60) != 1 {
		t.Errorf("Should contain one 60")
	}

	noNum1 := NoNum1()
	if noNum1.Count(50) != 0 {
		t.Errorf("Should not contain 50")
	}
}

func TestDistinct(t *testing.T) {
	someNum1 := SomeNum1(60)
	if someNum1.Distinct().Size() != 1 {
		t.Errorf("Should be length 1")
	}
	if someNum1.Distinct().Head() != 60 {
		t.Errorf("Should be length 60")
	}

	noNum1 := NoNum1()
	if noNum1.Distinct().NonEmpty() {
		t.Errorf("Should be empty")
	}
}

func TestMapToNum1(t *testing.T) {
	someNum1 := SomeNum1(60)
	fn1 := func(o Num1) Foo { return Foo(fmt.Sprintf("%d", o)) }
	m1 := someNum1.MapToFoo(fn1)
	if m1.Head() != "60" {
		t.Errorf("MapToFoo should be '60' but got %q", m1)
	}

	noNum1 := NoNum1()
	m2 := noNum1.MapToFoo(fn1)
	if m2.NonEmpty() {
		t.Errorf("MapToFoo should be absent but got %+v", m2)
	}
}

func TestFlatMapTo(t *testing.T) {
	someNum1 := SomeNum1(60)
	fn0 := func(o Num1) FooCollection { return NoFoo() }
	fn1 := func(o Num1) FooCollection { return SomeFoo(Foo(fmt.Sprintf("%d", o))) }

	m0 := someNum1.FlatMapToFoo(fn0)
	if m0.NonEmpty() {
		t.Errorf("MapToFoo should be absent but got %q", m0)
	}

	m1 := someNum1.FlatMapToFoo(fn1)
	if m1.Head() != "60" {
		t.Errorf("MapToFoo should be '60' but got %q", m1)
	}

	noNum1 := NoNum1()
	m2 := noNum1.FlatMapToFoo(fn1)
	if m2.NonEmpty() {
		t.Errorf("MapToFoo should be absent but got %+v", m2)
	}
}

func TestNum1MkString1(t *testing.T) {
	someNum1 := SomeNum1(123)

	s1 := someNum1.MkString3("<", ", ", ">")
	if s1 != "<123>" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := someNum1.MkString(", ")
	if s2 != "123" {
		t.Errorf("MkString got %q", s2)
	}

	s3 := someNum1.String()
	if s3 != "[123]" {
		t.Errorf("MkString got %q", s3)
	}
}

func TestNum1MkString1Empty(t *testing.T) {
	noNum1 := NoNum1()

	s1 := noNum1.MkString3("<", ", ", ">")
	if s1 != "<>" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := noNum1.MkString(", ")
	if s2 != "" {
		t.Errorf("MkString got %q", s2)
	}

	s3 := noNum1.String()
	if s3 != "[]" {
		t.Errorf("MkString got %q", s3)
	}
}

func TestFooMkString1(t *testing.T) {
	someFoo := SomeFoo("foo")

	s1 := someFoo.MkString3("<", ", ", ">")
	if s1 != "<foo>" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := someFoo.MkString(", ")
	if s2 != "foo" {
		t.Errorf("MkString got %q", s2)
	}

	s3 := someFoo.String()
	if s3 != "[foo]" {
		t.Errorf("MkString got %q", s3)
	}
}

func TestFooMkString1Empty(t *testing.T) {
	noFoo := NoFoo()

	s1 := noFoo.MkString3("<", ", ", ">")
	if s1 != "<>" {
		t.Errorf("MkString got %q", s1)
	}

	s2 := noFoo.MkString(", ")
	if s2 != "" {
		t.Errorf("MkString got %q", s2)
	}

	s3 := noFoo.String()
	if s3 != "[]" {
		t.Errorf("MkString got %q", s3)
	}
}
