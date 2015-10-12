// Generated by: setup
// TypeWriter: Option
// Directive: +test on Foo

package main

// OptionalFoo is an optional of type Foo. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalFoo interface {
	Get() Foo
	GetOrElse(d func() Foo) Foo
	OrElse(alternative func() OptionalFoo) OptionalFoo
	Len() int
	IsEmpty() bool
	NonEmpty() bool
	Find(fn func(Foo) bool) OptionalFoo
	Exists(fn func(Foo) bool) bool
	Forall(fn func(Foo) bool) bool
	Foreach(fn func(Foo))
	Filter(fn func(Foo) bool) (result OptionalFoo)
}

//-----------------------------------------------------------------------------

type SomeFoo Foo

// cross-check
var _ OptionalFoo = new(SomeFoo)

// Get returns the contained element if present. Otherwise it panics.
func (x SomeFoo) Get() Foo {
	return Foo(x)
}

// GetOrElse returns the contained element if present. Otherwise it returns a default.
func (x SomeFoo) GetOrElse(d func() Foo) Foo {
	return Foo(x)
}

// OrElse returns this option if present. Otherwise it returns the alternative.
func (x SomeFoo) OrElse(alternative func() OptionalFoo) OptionalFoo {
	return x
}

// Len returns 1
func (x SomeFoo) Len() int {
	return 1
}

// IsEmpty returns false.
func (x SomeFoo) IsEmpty() bool {
	return false
}

// NonEmpty returns true.
func (x SomeFoo) NonEmpty() bool {
	return true
}

// Find returns the contained value if it matches the predicate.
func (x SomeFoo) Find(predicate func(Foo) bool) OptionalFoo {
	if predicate(Foo(x)) {
		return x
	}
	return NoFoo()
}

// Exists verifies that one or more elements of SomeFoo return true for the passed predicate.
func (x SomeFoo) Exists(predicate func(Foo) bool) bool {
	return predicate(Foo(x))
}

// Forall verifies that all elements of OptionalFoo return true for the passed predicate.
func (x SomeFoo) Forall(predicate func(Foo) bool) bool {
	return predicate(Foo(x))
}

// Foreach iterates over OptionalFoo and executes the passed function against each element.
func (x SomeFoo) Foreach(fn func(Foo)) {
	fn(Foo(x))
}

// Filter returns a OptionalFoo whose elements return true for a predicate.
func (x SomeFoo) Filter(predicate func(Foo) bool) OptionalFoo {
	if predicate(Foo(x)) {
		return x
	}
	return NoFoo()
}

// These methods require Foo be comparable.

// Contains verifies that a given value is contained in OptionalFoo.
func (v SomeFoo) Contains(value Foo) bool {
	if Foo(v) == value {
		return true
	}
	return false
}

// Count gives the number elements of OptionalFoo that match a certain value.
func (v SomeFoo) Count(value Foo) (result int) {
	if Foo(v) == value {
		result++
	}
	return
}

//-----------------------------------------------------------------------------

type noFoo struct{}

// cross-check
var _ OptionalFoo = new(noFoo)

func NoFoo() OptionalFoo {
	return noFoo{}
}

// Get returns the contained element if present. Otherwise it panics.
func (x noFoo) Get() Foo {
	panic("Absent option - this indicates a promgramming error.")
}

// GetOrElse returns the contained element if present. Otherwise it returns a default.
func (x noFoo) GetOrElse(d func() Foo) Foo {
	return d()
}

// OrElse returns this option if present. Otherwise it returns the alternative.
func (x noFoo) OrElse(alternative func() OptionalFoo) OptionalFoo {
	return alternative()
}

// Len returns 0
func (x noFoo) Len() int {
	return 0
}

// IsEmpty returns true.
func (x noFoo) IsEmpty() bool {
	return true
}

// NonEmpty returns false.
func (x noFoo) NonEmpty() bool {
	return false
}

// Find can never succeed so returns NoFoo.
func (x noFoo) Find(predicate func(Foo) bool) OptionalFoo {
	return x
}

// Exists can never succeed so returns false.
func (x noFoo) Exists(predicate func(Foo) bool) bool {
	return false
}

// Forall verifies that all elements of OptionalFoo return true for the passed predicate, which
// is always true.
func (x noFoo) Forall(predicate func(Foo) bool) bool {
	return true
}

// Foreach iterates over OptionalFoo and executes the passed function against each element. There
// aren't any, so it does nothing.
func (x noFoo) Foreach(fn func(Foo)) {
	// does nothing
}

// Filter returns a OptionalFoo whose elements return true for a predicate.
func (x noFoo) Filter(predicate func(Foo) bool) (result OptionalFoo) {
	return x
}

// Contains verifies that a given value is contained in OptionalFoo.
func (v noFoo) Contains(value Foo) bool {
	return false
}

// Count gives the number elements of OptionalFoo that match a certain value.
func (v noFoo) Count(value Foo) int {
	return 0
}
