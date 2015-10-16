// Generated by: setup
// TypeWriter: Option
// Directive: +test on Foo

package main

// FooSeq is an interface for sequences of type Foo, including lists and options (where present).
type FooSeq interface {
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Foo

	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo))

	// Filter returns a new FooSeq whose elements return true for a predicate function.
	Filter(predicate func(Foo) bool) (result FooSeq)

	// Partition returns two new FooLists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Foo) bool) (matching FooSeq, others FooSeq)

	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Foo) bool) OptionalFoo

	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Foo is not comparable.
	Equals(other FooSeq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo is not comparable.
	Contains(value Foo) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo is not comparable.
	Count(value Foo) int
}

//-------------------------------------------------------------------------------------------------
// OptionalFoo is an optional of type Foo. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalFoo struct {
	x *Foo
}

// shared none value
var noneFoo = OptionalFoo{nil}

func NoFoo() OptionalFoo {
	return noneFoo
}

func SomeFoo(x Foo) OptionalFoo {

	return OptionalFoo{&x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalFoo) Head() Foo {

	return *(o.x)

}

func (o OptionalFoo) Get() Foo {
	return o.Head()
}

func (o OptionalFoo) GetOrElse(d func() Foo) Foo {
	if o.IsEmpty() {
		return d()
	}
	return *o.x
}

func (o OptionalFoo) OrElse(alternative func() OptionalFoo) OptionalFoo {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//----- FooSeq Methods -----

func (o OptionalFoo) Len() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalFoo) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalFoo) NonEmpty() bool {
	return o.x != nil
}

func (o OptionalFoo) Find(predicate func(Foo) bool) OptionalFoo {
	if o.IsEmpty() {
		return o
	}
	if predicate(*o.x) {
		return o
	}
	return noneFoo
}

func (o OptionalFoo) Exists(predicate func(Foo) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(*o.x)
}

func (o OptionalFoo) Forall(predicate func(Foo) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(*o.x)
}

func (o OptionalFoo) Foreach(fn func(Foo)) {
	if o.NonEmpty() {
		fn(*o.x)
	}
}

func (o OptionalFoo) Filter(predicate func(Foo) bool) FooSeq {
	return o.Find(predicate)
}

func (o OptionalFoo) Partition(predicate func(Foo) bool) (FooSeq, FooSeq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneFoo
	}
	return noneFoo, o
}

// These methods require Foo be comparable.

// Equals verifies that one or more elements of FooList return true for the passed func.
func (o OptionalFoo) Equals(other FooSeq) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Len() > 1 {
		return false
	}
	a := o.Head()
	b := other.Head()
	return a == b
}

func (o OptionalFoo) Contains(value Foo) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == value
}

func (o OptionalFoo) Count(value Foo) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new FooSeq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if Foo is not comparable.
func (o OptionalFoo) Distinct() FooSeq {
	return o
}
