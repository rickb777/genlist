// Generated by: setup
// TypeWriter: List
// Directive: +test on Num1

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

//-------------------------------------------------------------------------------------------------

// Num1Collection is an interface for collections of type Num1, including sets, lists and options (where present).
type Num1Collection interface {
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
	Head() Num1

	// ToSlice returns a plain slice containing all the elements in the collection. This is useful for bespoke iteration etc.
	// For sequences, the order of the elements is simple and well defined.
	// For non-sequences (i.e. sets) the order of the elements is stable but not well defined. This means it will give
	// the same order each subsequent time it is used as it did the first time.
	ToSlice() []Num1

	// ToInts gets all the elements in a slice of the underlying type, []int.
	ToInts() []int

	// ToList gets all the elements in a List.
	ToList() Num1List

	// Send sends all elements along a channel of type Num1.
	// For sequences, the order of the elements is simple and well defined.
	// For non-sequences (i.e. sets) the order of the elements is stable but not well defined. This means it will give
	// the same order each subsequent time it is used as it did the first time.
	Send() <-chan Num1

	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(Num1) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(Num1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Num1))

	// Filter returns a new Num1Collection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(Num1) bool) (result Num1Collection)

	// Partition returns two new Num1Collections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Num1) bool) (matching Num1Collection, others Num1Collection)

	// Equals verifies that another Num1Collection has the same size and elements as this one. Also,
	// if the collection is a sequence, the order must be the same.
	// Omitted if Num1 is not comparable.
	Equals(other Num1Collection) bool

	// Contains tests whether a given value is present in the collection.
	// Omitted if Num1 is not comparable.
	Contains(value Num1) bool

	// Sum sums Num1 elements.
	// Omitted if Num1 is not numeric.
	Sum() Num1

	// Mean computes the arithmetic mean of all elements. Panics if the collection is empty.
	// Omitted if Num1 is not numeric.
	Mean() float64

	// Min returns the minimum value of Num1List. In the case of multiple items being equally minimal,
	// the first such element is returned. Panics if the collection is empty.
	Min() Num1

	// Max returns the maximum value of Num1List. In the case of multiple items being equally maximal,
	// the first such element is returned. Panics if the collection is empty.
	Max() Num1

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

// Num1List is a slice of type Num1. Use it where you would use []Num1.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Num1List []Num1

//-------------------------------------------------------------------------------------------------

// NewNum1List constructs a new list containing the supplied values, if any.
func NewNum1List(values ...Num1) Num1List {
	list := make(Num1List, len(values))
	for i, v := range values {
		list[i] = v
	}
	return list
}

// NewNum1ListFromInts constructs a new Num1List from a []int.
func NewNum1ListFromInts(values []int) Num1List {
	list := make(Num1List, len(values))
	for i, v := range values {
		list[i] = Num1(v)
	}
	return list
}

// BuildNum1ListFromChan constructs a new Num1List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildNum1ListFromChan(source <-chan Num1) Num1List {
	result := make(Num1List, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list Num1List) Head() Num1 {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list Num1List) Last() Num1 {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list Num1List) Tail() Num1Collection {
	return Num1List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list Num1List) Init() Num1Collection {
	return Num1List(list[:len(list)-1])
}

// IsEmpty tests whether Num1List is empty.
func (list Num1List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Num1List is empty.
func (list Num1List) NonEmpty() bool {
	return len(list) > 0
}

// IsSequence returns true for lists.
func (list Num1List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list Num1List) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// ToSlice gets all the list's elements in a plain slice. This is simply a type conversion and is hardly needed
// for lists, because the underlying type can be used directly also.
// It is part of the Num1Collection interface.
func (list Num1List) ToSlice() []Num1 {
	return []Num1(list)
}

// ToInts gets all the elements in a []int.
func (list Num1List) ToInts() []int {
	slice := make([]int, len(list))
	for i, v := range list {
		slice[i] = int(v)
	}
	return slice
}

// ToList simply returns the list in this case, but is part of the Collection interface.
func (list Num1List) ToList() Num1List {
	return list
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list Num1List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list Num1List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list Num1List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// These methods require that Num1 be ordered.

// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list Num1List) Less(i, j int) bool {
	return list[i] < list[j]
}

// Sort returns a new ordered Num1List.
func (list Num1List) Sort() Num1List {
	result := make(Num1List, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether Num1List is sorted.
func (list Num1List) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered Num1List.
func (list Num1List) SortDesc() Num1List {
	result := make(Num1List, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether Num1List is reverse-sorted.
func (list Num1List) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of Num1List return true for the passed func.
func (list Num1List) Exists(fn func(Num1) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Num1List return true for the passed func.
func (list Num1List) Forall(fn func(Num1) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Num1List and executes the passed func against each element.
func (list Num1List) Foreach(fn func(Num1)) {
	for _, v := range list {
		fn(v)
	}
}

// Send gets a channel that will send all the elements in order.
func (list Num1List) Send() <-chan Num1 {
	ch := make(chan Num1)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of Num1List with all elements in the reverse order.
func (list Num1List) Reverse() Num1List {
	numItems := len(list)
	result := make(Num1List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of Num1List, using a version of the Fisher-Yates shuffle.
func (list Num1List) Shuffle() Num1List {
	numItems := len(list)
	result := make(Num1List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a new Num1List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num1List) Take(n int) Num1List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new Num1List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num1List) Drop(n int) Num1List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new Num1List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num1List) TakeLast(n int) Num1List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new Num1List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num1List) DropLast(n int) Num1List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new Num1List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list Num1List) TakeWhile(p func(Num1) bool) (result Num1List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new Num1List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list Num1List) DropWhile(p func(Num1) bool) (result Num1List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new Num1List whose elements return true for func.
func (list Num1List) Filter(fn func(Num1) bool) Num1Collection {
	result := make(Num1List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new Num1Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list Num1List) Partition(p func(Num1) bool) (Num1Collection, Num1Collection) {
	matching := make(Num1List, 0, len(list)/2)
	others := make(Num1List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of Num1List that return true for the passed predicate.
func (list Num1List) CountBy(predicate func(Num1) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Num1List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list Num1List) MinBy(less func(Num1, Num1) bool) (result Num1) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// MaxBy returns an element of Num1List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list Num1List) MaxBy(less func(Num1, Num1) bool) (result Num1) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list[m], list[i]) {
			m = i
		}
	}
	result = list[m]
	return
}

// DistinctBy returns a new Num1List whose elements are unique, where equality is defined by a passed func.
func (list Num1List) DistinctBy(equal func(Num1, Num1) bool) (result Num1List) {
Outer:
	for _, v := range list {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list Num1List) IndexWhere(p func(Num1) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list Num1List) IndexWhere2(p func(Num1) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list Num1List) LastIndexWhere(p func(Num1) bool) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list Num1List) LastIndexWhere2(p func(Num1) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

// Equals verifies that another Num1Collection has the same size and elements as this one. Also,
// because this collection is a sequence, the order must be the same.
// Omitted if Num1 is not comparable.
func (list Num1List) Equals(other Num1Collection) bool {
	if len(list) != other.Size() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a Num1) {
		if eq {
			v := list[i]
			if v != a {
				eq = false
			}
			i++
		}
	})
	return eq
}

//-------------------------------------------------------------------------------------------------
// These methods are provided because Num1 is comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list Num1List) IndexOf(value Num1) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list Num1List) IndexOf2(value Num1, from int) int {
	for i, v := range list {
		if i >= from && v == value {
			return i
		}
	}
	return -1
}

// Contains verifies that a given value is contained in Num1List.
func (list Num1List) Contains(value Num1) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of Num1List that match a certain value.
func (list Num1List) Count(value Num1) (result int) {
	for _, v := range list {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new Num1List whose elements are unique, retaining the original order.
func (list Num1List) Distinct() Num1Collection {
	result := make(Num1List, 0)
	appended := make(map[Num1]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

//-------------------------------------------------------------------------------------------------
// These methods are provided because Num1 is numeric.

// Sum sums all elements in the list.
func (list Num1List) Sum() (result Num1) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (list Num1List) Mean() float64 {
	l := len(list)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length Num1List")
	}
	sum := list.Sum()
	return float64(sum) / float64(l)
}

//-------------------------------------------------------------------------------------------------
// These methods are provided because Num1 is ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (list Num1List) Min() (result Num1) {
	if len(list) == 0 {
		panic("Cannot determine the Min of an empty list.")
	}
	result = list[0]
	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the element with the maximum value. In the case of multiple items being equally maximal,
// the first such element is returned. Panics if the collection is empty.
func (list Num1List) Max() (result Num1) {
	if len(list) == 0 {
		panic("Cannot determine the Max of an empty list.")
	}
	result = list[0]
	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list Num1List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list Num1List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list Num1List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}

//-------------------------------------------------------------------------------------------------
// List:MapTo[Foo]

// MapToFoo transforms Num1List to FooList.
func (list Num1List) MapToFoo(fn func(Num1) Foo) FooCollection {
	result := make(FooList, 0, len(list))
	for _, v := range list {
		u := fn(v)
		result = append(result, u)
	}
	return result
}

// FlatMapToFoo transforms Num1List to FooList, by repeatedly
// calling the supplied function and concatenating the results as a single flat list.
func (list Num1List) FlatMapToFoo(fn func(Num1) FooCollection) FooCollection {
	result := make(FooList, 0, len(list))
	for _, v := range list {
		u := fn(v)
		if u.NonEmpty() {
			result = append(result, (u.ToList())...)
		}
	}
	return result
}

// List flags: {Collection:false List:true Option:false Set:false Plumbing:false Tag:map[MapTo:true]}
