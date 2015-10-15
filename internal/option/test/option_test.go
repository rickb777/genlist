package main

import (
	"testing"
	"fmt"
)

func TestSomeOther(t *testing.T) {
	someThing := SomeOther(60)

	if someThing.Len() != 1 {
		t.Errorf("Len should be 1")
	}

	if someThing.IsEmpty() {
		t.Errorf("IsEmpty should be false")
	}

	if !someThing.NonEmpty() {
		t.Errorf("NonEmpty should be true")
	}

	if someThing.Get() != 60 {
		t.Errorf("Get should be 60")
	}

	if someThing.GetOrElse(func() Other { return 50 }) != 60 {
		t.Errorf("GetOrElse should be 60")
	}

	if someThing.OrElse(func() OptionalOther { return NoOther() }).Get() != 60 {
		t.Errorf("OrElse should be 60")
	}

	if !someThing.Exists(func(Other) bool { return true }) {
		t.Errorf("Exists should be true")
	}

	if !someThing.Forall(func(Other) bool { return true }) {
		t.Errorf("Forall should be true")
	}

	if someThing.Filter(func(Other) bool { return true }).(OptionalOther).Get() != 60 {
		t.Errorf("Filter should be 60")
	}

	if someThing.Find(func(Other) bool { return true }).Get() != 60 {
		t.Errorf("Find should be 60")
	}

	if someThing.Filter(func(Other) bool { return false }).NonEmpty() {
		t.Errorf("Filter should be empty")
	}

	x := Other(0)
	someThing.Foreach(func(Other) { x = Other(1) })
	if x != 1 {
		t.Errorf("Foreach should set x")
	}

	if someThing.Contains(50) {
		t.Errorf("Should not contain 50")
	}

	if !someThing.Contains(60) {
		t.Errorf("Should contain 60")
	}

	if someThing.Count(50) != 0 {
		t.Errorf("Should contain zero 50s")
	}

	if someThing.Count(60) != 1 {
		t.Errorf("Should contain one 60")
	}

	if someThing.Sum() != 60 {
		t.Errorf("Sum should be 60")
	}

	m := someThing.MapToFoo(func(o Other) Foo { return Foo(fmt.Sprintf("%d", o)) }).Get()
	if m != "60" {
		t.Errorf("MapToFoo should be '60' but got %q", m)
	}
}

func TestNoOther(t *testing.T) {
	noThing := NoOther()

	if noThing.Len() != 0 {
		t.Errorf("Len should be 0")
	}

	if !noThing.IsEmpty() {
		t.Errorf("IsEmpty should be true")
	}

	if noThing.NonEmpty() {
		t.Errorf("NonEmpty should be false")
	}

	if noThing.GetOrElse(func() Other { return 50 }) != 50 {
		t.Errorf("GetOrElse should be 50")
	}

	if noThing.OrElse(func() OptionalOther { return SomeOther(60) }).Get() != 60 {
		t.Errorf("OrElse should be 60")
	}

	if noThing.Exists(func(Other) bool { return true }) {
		t.Errorf("Exists should be false")
	}

	if !noThing.Forall(func(Other) bool { return true }) {
		t.Errorf("Forall should be true")
	}

	if noThing.Filter(func(Other) bool { return true }).NonEmpty() {
		t.Errorf("Filter should be empty")
	}

	if noThing.Find(func(Other) bool { return true }).NonEmpty() {
		t.Errorf("Find should be empty")
	}

	x := Other(0)
	noThing.Foreach(func(Other) { x = Other(1) })
	if x != 0 {
		t.Errorf("Foreach should not set x")
	}

	if noThing.Contains(50) {
		t.Errorf("Should not contain 50")
	}

	if noThing.Count(50) != 0 {
		t.Errorf("Should not contain 50")
	}

	if noThing.Sum() != 0 {
		t.Errorf("Sum should be 0")
	}
}
