// Generated by: setup
// TypeWriter: List
// Directive: +test on Num1

package main

import (
	"errors"
	"math/rand"
	"sort"
)

// Num1Seq is an interface for sequences of type Num1, including lists and options (where present).
type Num1Seq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Num1) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Num1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Num1))

	// Filter returns a new Num1Seq whose elements return true for func.
	Filter(predicate func(Num1) bool) (result Num1Seq)

	// Converts the sequence to a list. For lists, this is merely a type conversion.
	ToList() Num1List

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Num1 is not comparable.
	Contains(value Num1) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Num1 is not comparable.
	Count(value Num1) int

	// Distinct returns a new Num1Seq whose elements are all unique.
	// Omitted if Num1 is not comparable.
	Distinct() Num1Seq

	// Sum sums Num1 elements.
	// Omitted if Num1 is not numeric.
	Sum() Num1
}

//-------------------------------------------------------------------------------------------------
// Num1List is a slice of type Num1. Use it where you would use []Num1.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Num1List []Num1

//-------------------------------------------------------------------------------------------------

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
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

// IsEmpty tests whether Num1List is empty.
func (list Num1List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Num1List is empty.
func (list Num1List) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list Num1List) ToList() Num1List {
	return list
}

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

// Filter returns a new Num1List whose elements return true for func.
func (list Num1List) Filter(fn func(Num1) bool) Num1Seq {
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
func (list Num1List) Partition(p func(Num1) bool) (matching Num1List, others Num1List) {
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
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
// element is returned. Returns error if no elements.
func (list Num1List) MinBy(less func(Num1, Num1) bool) (result Num1, err error) {
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

// MaxBy returns an element of Num1List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list Num1List) MaxBy(less func(Num1, Num1) bool) (result Num1, err error) {
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

// These methods require Num1 be comparable.

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

// Distinct returns a new Num1List whose elements are unique.
func (list Num1List) Distinct() Num1Seq {
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

// These methods require Num1 be numeric.

// Sum sums Num1 elements in Num1List.
func (list Num1List) Sum() (result Num1) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean sums Num1List over all elements and divides by len(Num1List).
func (list Num1List) Mean() (Num1, error) {
	var result Num1

	l := len(list)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length Num1List")
	}
	for _, v := range list {
		result += v
	}
	result = result / Num1(l)
	return result, nil
}

// These methods require Num1 be ordered.

// Min returns the minimum value of Num1List. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements.
func (list Num1List) Min() (result Num1, err error) {
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

// Max returns the maximum value of Num1List. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements.
func (list Num1List) Max() (result Num1, err error) {
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

// optionForList
