// Generated by: setup
// TypeWriter: Option
// Directive: +test on *Bar

package main

import "fmt"

//-------------------------------------------------------------------------------------------------
// BarCollection is an interface for collections of type Bar, including sets, lists and options (where present).
type BarCollection interface {
	// Size gets the size/length of the collection.
	Size() int

	// IsEmpty returns true if the collection is empty.
	IsEmpty() bool

	// NonEmpty returns true if the collection is non-empty.
	NonEmpty() bool

	// IsSequence returns true for lists, but false otherwise.
	IsSequence() bool

	// IsSet returns true for sets, but false otherwise.
	IsSet() bool

	// Head returns the first element of a list or an arbitrary element of a set or the contents of an option.
	// Panics if the collection is empty.
	Head() *Bar

	//-------------------------------------------------------------------------
	// ToSlice returns a plain slice containing all the elements in the collection.
	// This is useful for bespoke iteration etc.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	ToSlice() []*Bar

	// Send sends all elements along a channel of type Bar.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	Send() <-chan *Bar

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(*Bar) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(*Bar) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(*Bar))

	//-------------------------------------------------------------------------
	// Filter returns a new BarCollection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(*Bar) bool) (result BarCollection)

	// Partition returns two new BarCollections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(*Bar) bool) (matching BarCollection, others BarCollection)

	//-------------------------------------------------------------------------
	// Equals verifies that another BarCollection has the same size and elements as this one. Also,
	// if the collection is a sequence, the order must be the same.
	// Omitted if Bar is not comparable.
	Equals(other BarCollection) bool

	// Contains tests whether a given value is present in the collection.
	// Omitted if Bar is not comparable.
	Contains(value *Bar) bool

	// Min returns an element of BarList containing the minimum value, when compared to other elements
	// using a specified comparator function defining ‘less’. For ordered sequences, Min returns the first such element.
	// Panics if the collection is empty.
	Min(less func(*Bar, *Bar) bool) *Bar

	// Max returns an element of BarList containing the maximum value, when compared to other elements
	// using a specified comparator function defining ‘less’. For ordered sequences, Max returns the first such element.
	// Panics if the collection is empty.
	Max(less func(*Bar, *Bar) bool) *Bar

	//-------------------------------------------------------------------------
	// String gets a string representation of the collection. "[" and "]" surround
	// a comma-separated list of the elements.
	String() string

	// MkString gets a string representation of the collection. "[" and "]" surround a list
	// of the elements joined by the separator you provide.
	MkString(sep string) string

	// MkString3 gets a string representation of the collection. 'pfx' and 'sfx' surround a list
	// of the elements joined by the 'mid' separator you provide.
	MkString3(pfx, mid, sfx string) string
}

//-------------------------------------------------------------------------------------------------

// OptionalBar is an optional of type *Bar. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options. In particular, an option is a collection
// with a maximum cardinality of one. As such, options can be converted to/from lists and sets.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option
type OptionalBar struct {
	x *Bar
}

// shared none value
var noneBar = OptionalBar{nil}

// NoBar gets an empty instance.
func NoBar() OptionalBar {
	return noneBar
}

// SomeBar gets a non-empty instance wrapping some value *x*.
func SomeBar(x *Bar) OptionalBar {

	if x == nil {
		return NoBar()
	}
	return OptionalBar{x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalBar) Head() *Bar {
	return o.Get()
}

func (o OptionalBar) Get() *Bar {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return o.x
}

func (o OptionalBar) GetOrElse(d func() *Bar) *Bar {
	if o.IsEmpty() {
		return d()
	}
	return o.x
}

func (o OptionalBar) OrElse(alternative func() OptionalBar) OptionalBar {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalBar) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalBar) Len() int {
	return o.Size()
}

func (o OptionalBar) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalBar) NonEmpty() bool {
	return o.x != nil
}

// IsSequence returns false for options.
func (o OptionalBar) IsSequence() bool {
	return false
}

// IsSet returns false for options.
func (o OptionalBar) IsSet() bool {
	return false
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalBar) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalBar) Find(predicate func(*Bar) bool) OptionalBar {
	if o.IsEmpty() {
		return o
	}
	if predicate(o.x) {
		return o
	}
	return noneBar
}

func (o OptionalBar) Exists(predicate func(*Bar) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(o.x)
}

func (o OptionalBar) Forall(predicate func(*Bar) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(o.x)
}

func (o OptionalBar) Foreach(fn func(*Bar)) {
	if o.NonEmpty() {
		fn(o.x)
	}
}

// Send gets a channel that will send all the elements in order.
func (o OptionalBar) Send() <-chan *Bar {
	ch := make(chan *Bar)
	go func() {
		if o.NonEmpty() {
			ch <- o.x
		}
		close(ch)
	}()
	return ch
}

func (o OptionalBar) Filter(predicate func(*Bar) bool) BarCollection {
	return o.Find(predicate)
}

func (o OptionalBar) Partition(predicate func(*Bar) bool) (BarCollection, BarCollection) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(o.x) {
		return o, noneBar
	}
	return noneBar, o
}

func (o OptionalBar) ToSlice() []*Bar {
	slice := make([]*Bar, o.Size())
	if o.NonEmpty() {
		slice[0] = o.x
	}
	return slice
}

//-------------------------------------------------------------------------------------------------
// These methods require *Bar be comparable.

// Equals verifies that one or more elements of BarList return true for the passed func.
func (o OptionalBar) Equals(other BarCollection) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Size() > 1 {
		return false
	}
	a := o.x
	s := other.ToSlice()
	b := s[0]
	return *a == *b
}

func (o OptionalBar) Contains(value *Bar) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == *value
}

func (o OptionalBar) Count(value *Bar) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new BarCollection whose elements are all unique. For options, this simply returns the
// receiver.
// Omitted if Bar is not comparable.
func (o OptionalBar) Distinct() BarCollection {
	return o
}

// Min returns an element of BarList containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’. For ordered sequences, Min returns the first such element.
// Panics if the collection is empty.
func (o OptionalBar) Min(less func(*Bar, *Bar) bool) *Bar {
	return o.Get()
}

// Max returns an element of BarList containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’. For ordered sequences, Max returns the first such element.
// Panics if the collection is empty.
func (o OptionalBar) Max(less func(*Bar, *Bar) bool) *Bar {
	return o.Get()
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the option as an array of one element.
func (o OptionalBar) String() string {
	return o.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (o OptionalBar) MkString(sep string) string {
	return o.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (o OptionalBar) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

// Option flags: {Collection:false List:false Option:true Set:false Plumbing:false Tag:map[]}
