// Generated by: setup
// TypeWriter: List
// Directive: +test on Foo

package main

import (
	"errors"
	"math/rand"
	"sort"
)

// FooList is a slice of type Foo. Use it where you would use []Foo.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type FooList []Foo

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// These methods require Foo be ordered.

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

// IsEmpty tests whether FooList is empty.
func (list FooList) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether FooList is empty.
func (list FooList) NonEmpty() bool {
	return len(list) > 0
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

// Shuffle returns a shuffled copy of FooList, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
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

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list FooList) ToList() FooList {
	return list
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
func (list FooList) Filter(fn func(Foo) bool) (result FooList) {
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
func (list FooList) Partition(p func(Foo) bool) (matching FooList, others FooList) {
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
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

// DistinctBy returns a new FooList whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
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

// These methods require Foo be comarable.

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

// Distinct returns a new FooList whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (list FooList) Distinct() (result FooList) {
	appended := make(map[Foo]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// These methods require Foo be numeric.

// Sum sums Foo elements in FooList. See: http://clipperhouse.github.io/gen/#Sum
func (list FooList) Sum() (result Foo) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean sums FooList over all elements and divides by len(FooList). See: http://clipperhouse.github.io/gen/#Mean
func (list FooList) Mean() (Foo, error) {
	var result Foo

	l := len(list)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length FooList")
	}
	for _, v := range list {
		result += v
	}
	result = result / Foo(l)
	return result, nil
}

// These methods require Foo be ordered.

// Min returns the minimum value of FooList. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
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
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
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
