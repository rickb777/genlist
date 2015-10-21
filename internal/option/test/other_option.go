// Generated by: setup
// TypeWriter: Option
// Directive: +test on Other

package main

import "fmt"

// OtherCollection is an interface for collections of type Other, including sets, lists and options (where present).
type OtherCollection interface {
	// Size gets the size/length of the sequence.
	Size() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool
}

// OtherSeq is an interface for sequences of type Other, including lists and options (where present).
type OtherSeq interface {
	OtherCollection
	// Len gets the size/length of the sequence - an alias for Size()
	Len() int

	//-------------------------------------------------------------------------
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Other

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() Other

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() OtherSeq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() OtherSeq

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Other) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Other) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Other))

	//-------------------------------------------------------------------------
	// Filter returns a new OtherSeq whose elements return true for a predicate function.
	Filter(predicate func(Other) bool) (result OtherSeq)

	// Partition returns two new OtherLists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Other) bool) (matching OtherSeq, others OtherSeq)

	//-------------------------------------------------------------------------
	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Other) bool) OptionalOther

	//-------------------------------------------------------------------------
	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Other is not comparable.
	Equals(other OtherSeq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Other is not comparable.
	Contains(value Other) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Other is not comparable.
	Count(value Other) int

	//-------------------------------------------------------------------------
	// Sum sums Other elements.
	// Omitted if Other is not numeric.
	Sum() Other

	// Mean computes the arithmetic mean of all elements.
	// Panics if the list is empty.
	Mean() Other
}

//-------------------------------------------------------------------------------------------------
// OptionalOther is an optional of type Other. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalOther struct {
	x *Other
}

// shared none value
var noneOther = OptionalOther{nil}

func NoOther() OptionalOther {
	return noneOther
}

func SomeOther(x Other) OptionalOther {

	return OptionalOther{&x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalOther) Head() Other {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return *(o.x)
}

// panics if option is empty
func (o OptionalOther) Last() Other {
	return o.Head()
}

// panics if option is empty
func (o OptionalOther) Tail() OtherSeq {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return noneOther
}

// panics if option is empty
func (o OptionalOther) Init() OtherSeq {
	return o.Tail()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalOther) Get() Other {
	return o.Head()
}

func (o OptionalOther) GetOrElse(d func() Other) Other {
	if o.IsEmpty() {
		return d()
	}
	return *o.x
}

func (o OptionalOther) OrElse(alternative func() OptionalOther) OptionalOther {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalOther) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalOther) Len() int {
	return o.Size()
}

func (o OptionalOther) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalOther) NonEmpty() bool {
	return o.x != nil
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalOther) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalOther) Find(predicate func(Other) bool) OptionalOther {
	if o.IsEmpty() {
		return o
	}
	if predicate(*o.x) {
		return o
	}
	return noneOther
}

func (o OptionalOther) Exists(predicate func(Other) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(*o.x)
}

func (o OptionalOther) Forall(predicate func(Other) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(*o.x)
}

func (o OptionalOther) Foreach(fn func(Other)) {
	if o.NonEmpty() {
		fn(*o.x)
	}
}

func (o OptionalOther) Filter(predicate func(Other) bool) OtherSeq {
	return o.Find(predicate)
}

func (o OptionalOther) Partition(predicate func(Other) bool) (OtherSeq, OtherSeq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneOther
	}
	return noneOther, o
}

//-------------------------------------------------------------------------------------------------
// These methods require Other be comparable.

// Equals verifies that one or more elements of OtherList return true for the passed func.
func (o OptionalOther) Equals(other OtherSeq) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Size() > 1 {
		return false
	}
	a := o.Head()
	b := other.Head()
	return a == b
}

func (o OptionalOther) Contains(value Other) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == value
}

func (o OptionalOther) Count(value Other) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new OtherSeq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if Other is not comparable.
func (o OptionalOther) Distinct() OtherSeq {
	return o
}

//-------------------------------------------------------------------------------------------------
// Sum sums Other elements.
// Omitted if Other is not numeric.
func (o OptionalOther) Sum() Other {

	if o.IsEmpty() {
		return 0
	}
	return *(o.x)

}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (o OptionalOther) Mean() Other {
	if o.IsEmpty() {
		panic("Cannot compute the arithmetic mean of zero-length OptionalOther")
	}
	return o.Sum()
}

//-------------------------------------------------------------------------------------------------
// String implements the Stringer interface to render the option as an array of one element.
func (o OptionalOther) String() string {
	return o.MkString(",")
}

// MkString concatenates the values as a string.
func (o OptionalOther) MkString(sep string) string {
	return o.MkString3("[", sep, "]")
}

// MkString3 concatenates the values as a string.
func (o OptionalOther) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

// MapToFoo transforms OptionalOther to OptionalFoo.
func (o OptionalOther) MapToFoo(fn func(Other) Foo) FooSeq {
	if o.IsEmpty() {
		return NoFoo()
	}

	u := fn(*(o.x))

	return SomeFoo(u)
}

// FlatMapToFoo transforms OptionalOther to OptionalFoo, by
// calling the supplied function on the enclosed instance, if any, and returning an option.
// The result is only defined if *both* the receiver is defined and the function returns a non-empty sequence.
func (o OptionalOther) FlatMapToFoo(fn func(Other) FooSeq) (result FooSeq) {
	if o.IsEmpty() {
		return NoFoo()
	}
	u := fn(*(o.x))
	if u.IsEmpty() {
		return NoFoo()
	}
	return SomeFoo(u.Head())
}

// Option flags: {Collection:false Sequence:false List:false Option:true Set:false Tag:map[With:true]}
