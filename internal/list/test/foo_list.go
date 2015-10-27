// Generated by: setup
// TypeWriter: List
// Directive: +test on Foo

package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

//-------------------------------------------------------------------------------------------------
// FooCollection is an interface for collections of type Foo, including sets, lists and options (where present).
type FooCollection interface {
	// Size gets the size/length of the sequence.
	Size() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo))

	// Iter sends all elements along a channel of type Foo.
	// The first time it is used, order of the elements is not well defined. But the order is stable, which means
	// it will give the same order each subsequent time it is used.
	Iter() <-chan Foo

	//-------------------------------------------------------------------------
	// Filter returns a new FooCollection whose elements return true for a predicate function.
	Filter(predicate func(Foo) bool) (result FooCollection)

	// Partition returns two new FooCollections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Foo) bool) (matching FooCollection, others FooCollection)

	//-------------------------------------------------------------------------
	// These methods require Foo be comparable.

	// Equals verifies that one or more elements of FooCollection return true for the passed func.
	Equals(other FooCollection) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo is not comparable.
	Contains(value Foo) bool
}

//-------------------------------------------------------------------------------------------------
// FooSeq is an interface for sequences of type Foo, including lists and options (where present).
type FooSeq interface {
	FooCollection

	// Len gets the size/length of the sequence - an alias for Size()
	Len() int

	//-------------------------------------------------------------------------
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Foo

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() Foo

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() FooSeq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() FooSeq

	// Converts the sequence to a list. For lists, this is merely a type assertion.
	ToList() FooList

	//-------------------------------------------------------------------------
	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo is not comparable.
	Count(value Foo) int

	// Distinct returns a new FooSeq whose elements are all unique.
	// Omitted if Foo is not comparable.
	Distinct() FooSeq
}

//-------------------------------------------------------------------------------------------------
// FooList is a slice of type Foo. Use it where you would use []Foo.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type FooList []Foo

//-------------------------------------------------------------------------------------------------
// BuildFooListFrom constructs a new FooList from a channel that supplies values
// until it is closed.
func BuildFooListFrom(source <-chan Foo) FooList {
	result := make(FooList, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// panics if list is empty
func (list FooList) Head() Foo {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// panics if list is empty
func (list FooList) Last() Foo {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// panics if list is empty
func (list FooList) Tail() FooSeq {
	return FooList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// panics if list is empty
func (list FooList) Init() FooSeq {
	return FooList(list[:len(list)-1])
}

// IsEmpty tests whether FooList is empty.
func (list FooList) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether FooList is empty.
func (list FooList) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list FooList) ToList() FooList {
	return list
}

// Size returns the number of items in the list - an alias of Len().
func (list FooList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// These methods require that Foo be ordered.

// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Less(i, j int) bool {
	return list[i] < list[j]
}

// Sort returns a new ordered FooList.
func (list FooList) Sort() FooList {
	result := make(FooList, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether FooList is sorted.
func (list FooList) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered FooList.
func (list FooList) SortDesc() FooList {
	result := make(FooList, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether FooList is reverse-sorted.
func (list FooList) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

// Exists verifies that one or more elements of FooList return true for the passed func.
func (list FooList) Exists(fn func(Foo) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of FooList return true for the passed func.
func (list FooList) Forall(fn func(Foo) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over FooList and executes the passed func against each element.
func (list FooList) Foreach(fn func(Foo)) {
	for _, v := range list {
		fn(v)
	}
}

// Iter gets a channel that will send all the elements in order.
func (list FooList) Iter() <-chan Foo {
	ch := make(chan Foo)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of FooList with all elements in the reverse order.
func (list FooList) Reverse() FooList {
	numItems := len(list)
	result := make(FooList, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of FooList, using a version of the Fisher-Yates shuffle.
func (list FooList) Shuffle() FooList {
	numItems := len(list)
	result := make(FooList, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

// Take returns a new FooList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) Take(n int) FooList {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new FooList without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) Drop(n int) FooList {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new FooList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) TakeLast(n int) FooList {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new FooList without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) DropLast(n int) FooList {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new FooList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list FooList) TakeWhile(p func(Foo) bool) (result FooList) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new FooList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list FooList) DropWhile(p func(Foo) bool) (result FooList) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Filter returns a new FooList whose elements return true for func.
func (list FooList) Filter(fn func(Foo) bool) FooCollection {
	result := make(FooList, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new FooLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list FooList) Partition(p func(Foo) bool) (FooCollection, FooCollection) {
	matching := make(FooList, 0, len(list)/2)
	others := make(FooList, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of FooList that return true for the passed predicate.
func (list FooList) CountBy(predicate func(Foo) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of FooList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list FooList) MinBy(less func(Foo, Foo) bool) (result Foo, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the MinBy of an empty list.")
		return
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

// MaxBy returns an element of FooList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list FooList) MaxBy(less func(Foo, Foo) bool) (result Foo, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the MaxBy of an empty list.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if list[i] != list[m] && !less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// DistinctBy returns a new FooList whose elements are unique, where equality is defined by a passed func.
func (list FooList) DistinctBy(equal func(Foo, Foo) bool) (result FooList) {
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
func (list FooList) IndexWhere(p func(Foo) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list FooList) IndexWhere2(p func(Foo) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list FooList) LastIndexWhere(p func(Foo) bool) int {
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
func (list FooList) LastIndexWhere2(p func(Foo) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

// These methods require Foo be comparable.

// Equals verifies that one or more elements of FooList return true for the passed func.
func (list FooList) Equals(other FooCollection) bool {
	if len(list) != other.Size() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a Foo) {
		if eq {
			v := list[i]
			if v != a {
				eq = false
			}
			i += 1
		}
	})
	return eq
}

// These methods require Foo be comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list FooList) IndexOf(value Foo) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list FooList) IndexOf2(value Foo, from int) int {
	for i, v := range list {
		if i >= from && v == value {
			return i
		}
	}
	return -1
}

// Contains verifies that a given value is contained in FooList.
func (list FooList) Contains(value Foo) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of FooList that match a certain value.
func (list FooList) Count(value Foo) (result int) {
	for _, v := range list {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new FooList whose elements are unique.
func (list FooList) Distinct() FooSeq {
	result := make(FooList, 0)
	appended := make(map[Foo]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// These methods require Foo be ordered.

// Min returns the minimum value of FooList. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements.
func (list FooList) Min() (result Foo, err error) {
	if len(list) == 0 {
		err = errors.New("Cannot determine the Min of an empty list.")
		return
	}
	result = list[0]
	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the maximum value of FooList. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements.
func (list FooList) Max() (result Foo, err error) {
	if len(list) == 0 {
		err = errors.New("Cannot determine the Max of an empty list.")
		return
	}
	result = list[0]
	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return
}

// String implements the Stringer interface to render the list as a comma-separated array.
func (list FooList) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (list FooList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (list FooList) MkString3(pfx, mid, sfx string) string {
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

// optionForList

// MapToNum1 transforms FooList to Num1List.
func (list FooList) MapToNum1(fn func(Foo) Num1) Num1Seq {
	result := make(Num1List, 0, len(list))
	for _, v := range list {
		u := fn(v)
		result = append(result, u)
	}
	return result
}

// FlatMapToNum1 transforms FooList to Num1List, by repeatedly
// calling the supplied function and concatenating the results as a single flat list.
func (list FooList) FlatMapToNum1(fn func(Foo) Num1Seq) Num1Seq {
	result := make(Num1List, 0, len(list))
	for _, v := range list {
		u := fn(v)
		if u.NonEmpty() {
			result = append(result, (u.ToList())...)
		}
	}
	return result
}

// List flags: {Collection:false Sequence:false List:true Option:false Set:false Tag:map[MapTo:true]}
