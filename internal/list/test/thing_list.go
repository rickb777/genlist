// Generated by: setup
// TypeWriter: List
// Directive: +test on Thing

package main

import (
	"errors"
	"math/rand"
)

// ThingSeq is an interface for sequences of type Thing, including lists and options (where present).
type ThingSeq interface {
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Thing

	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Thing) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Thing) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Thing))

	// Filter returns a new ThingSeq whose elements return true for a predicate function.
	Filter(predicate func(Thing) bool) (result ThingSeq)

	// Partition returns two new ThingLists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Thing) bool) (matching ThingSeq, others ThingSeq)

	// Converts the sequence to a list. For lists, this is merely a type conversion.
	ToList() ThingList

	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Thing is not comparable.
	Equals(other ThingSeq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Thing is not comparable.
	Contains(value Thing) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Thing is not comparable.
	Count(value Thing) int

	// Distinct returns a new ThingSeq whose elements are all unique.
	// Omitted if Thing is not comparable.
	Distinct() ThingSeq
}

//-------------------------------------------------------------------------------------------------
// ThingList is a slice of type Thing. Use it where you would use []Thing.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type ThingList []Thing

//-------------------------------------------------------------------------------------------------

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list ThingList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list ThingList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// panics if list is empty
func (list ThingList) Head() Thing {
	return list[0]
}

// IsEmpty tests whether ThingList is empty.
func (list ThingList) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether ThingList is empty.
func (list ThingList) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list ThingList) ToList() ThingList {
	return list
}

// Exists verifies that one or more elements of ThingList return true for the passed func.
func (list ThingList) Exists(fn func(Thing) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of ThingList return true for the passed func.
func (list ThingList) Forall(fn func(Thing) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over ThingList and executes the passed func against each element.
func (list ThingList) Foreach(fn func(Thing)) {
	for _, v := range list {
		fn(v)
	}
}

// Reverse returns a copy of ThingList with all elements in the reverse order.
func (list ThingList) Reverse() ThingList {
	numItems := len(list)
	result := make(ThingList, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of ThingList, using a version of the Fisher-Yates shuffle.
func (list ThingList) Shuffle() ThingList {
	numItems := len(list)
	result := make(ThingList, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

// Take returns a new ThingList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list ThingList) Take(n int) ThingList {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new ThingList without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list ThingList) Drop(n int) ThingList {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new ThingList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list ThingList) TakeLast(n int) ThingList {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new ThingList without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list ThingList) DropLast(n int) ThingList {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new ThingList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list ThingList) TakeWhile(p func(Thing) bool) (result ThingList) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new ThingList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list ThingList) DropWhile(p func(Thing) bool) (result ThingList) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Filter returns a new ThingList whose elements return true for func.
func (list ThingList) Filter(fn func(Thing) bool) ThingSeq {
	result := make(ThingList, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new ThingLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list ThingList) Partition(p func(Thing) bool) (ThingSeq, ThingSeq) {
	matching := make(ThingList, 0, len(list)/2)
	others := make(ThingList, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of ThingList that return true for the passed predicate.
func (list ThingList) CountBy(predicate func(Thing) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of ThingList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list ThingList) MinBy(less func(Thing, Thing) bool) (result Thing, err error) {
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

// MaxBy returns an element of ThingList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list ThingList) MaxBy(less func(Thing, Thing) bool) (result Thing, err error) {
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

// DistinctBy returns a new ThingList whose elements are unique, where equality is defined by a passed func.
func (list ThingList) DistinctBy(equal func(Thing, Thing) bool) (result ThingList) {
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
func (list ThingList) IndexWhere(p func(Thing) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate after or at some start index.
// If none exists, -1 is returned.
func (list ThingList) IndexWhere2(p func(Thing) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// These methods require Thing be comparable.

// Equals verifies that one or more elements of ThingList return true for the passed func.
func (list ThingList) Equals(other ThingSeq) bool {
	if len(list) != other.Len() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a Thing) {
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

// These methods require Thing be comparable.

// Contains verifies that a given value is contained in ThingList.
func (list ThingList) Contains(value Thing) bool {
	for _, v := range list {

		if v == value {
			return true
		}

	}
	return false
}

// Count gives the number elements of ThingList that match a certain value.
func (list ThingList) Count(value Thing) (result int) {
	for _, v := range list {

		if v == value {
			result++
		}

	}
	return
}

// Distinct returns a new ThingList whose elements are unique.
func (list ThingList) Distinct() ThingSeq {
	result := make(ThingList, 0)
	appended := make(map[Thing]bool)
	for _, v := range list {

		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}

	}
	return result
}

// Min returns the first element of ThingList containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Returns an error if the ThingList is empty.
func (list ThingList) Min(less func(Thing, Thing) bool) (result Thing, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the minimum of an empty list.")
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

// Max returns the first element of ThingList containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Returns an error if the ThingList is empty.
func (list ThingList) Max(less func(Thing, Thing) bool) (result Thing, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the maximum of an empty list.")
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

// optionForList

// MapToNum1 transforms ThingList to Num1List.
func (list ThingList) MapToNum1(fn func(Thing) Num1) (result Num1List) {
	for _, v := range list {

		result = append(result, fn(v))

	}
	return
}

// MapToNum2 transforms ThingList to Num2List.
func (list ThingList) MapToNum2(fn func(Thing) *Num2) (result Num2List) {
	for _, v := range list {

		result = append(result, fn(v))

	}
	return
}

// FoldLeftNum1 applies a binary operator to a start value and all elements of this list, going left to right.
func (list ThingList) FoldLeftNum1(zero Num1, fn func(Num1, Thing) Num1) Num1 {
	sum := zero
	for _, v := range list {
		sum = fn(sum, v)
	}
	return sum
}

// FoldRightNum1 applies a binary operator to a start value and all elements of this list, going right to left.
func (list ThingList) FoldRightNum1(zero Num1, fn func(Num1, Thing) Num1) Num1 {
	sum := zero
	for i := len(list) - 1; i >= 0; i-- {
		sum = fn(sum, list[i])
	}
	return sum
}

// This methods require Thing be comparable.

// GroupByNum1 groups elements into a map keyed by Num1.
func (list ThingList) GroupByNum1(fn func(Thing) Num1) map[Num1]ThingList {
	result := make(map[Num1]ThingList)
	for _, v := range list {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// These methods require Thing be numeric.

// SumNum1 sums Thing over elements in ThingList.
func (list ThingList) SumNum1(fn func(Thing) Num1) (result Num1) {
	for _, v := range list {
		result += fn(v)
	}
	return
}

// MeanNum1 sums Num1 over all elements and divides by len(ThingList).
func (list ThingList) MeanNum1(fn func(Thing) Num1) (result Num1, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Mean[Num1] of zero-length ThingList")
		return
	}
	for _, v := range list {
		result += fn(v)
	}
	result = result / Num1(l)
	return
}

// These methods require Num1 be ordered.

// MinByNum1 finds the first element which yields the smallest value measured by function fn.
// fn is usually called a projection or measuring function.
// Returns an error if the ThingList is empty.
func (list ThingList) MinByNum1(fn func(Thing) Num1) (result Thing, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length ThingList")
		return
	}
	result = list[0]
	if l > 1 {
		min := fn(result)
		for i := 1; i < l; i++ {
			v := list[i]
			f := fn(v)
			if min > f {
				min = f
				result = v
			}
		}
	}
	return
}

// MaxByNum1 finds the first element which yields the largest value measured by function fn.
// fn is usually called a projection or measuring function.
// Returns an error if the ThingList is empty.
func (list ThingList) MaxByNum1(fn func(Thing) Num1) (result Thing, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length ThingList")
		return
	}
	result = list[0]
	if l > 1 {
		max := fn(result)
		for i := 1; i < l; i++ {
			v := list[i]
			f := fn(v)
			if max < f {
				max = f
				result = v
			}
		}
	}
	return
}

// FoldLeftColour applies a binary operator to a start value and all elements of this list, going left to right.
func (list ThingList) FoldLeftColour(zero Colour, fn func(Colour, Thing) Colour) Colour {
	sum := zero
	for _, v := range list {
		sum = fn(sum, v)
	}
	return sum
}

// FoldRightColour applies a binary operator to a start value and all elements of this list, going right to left.
func (list ThingList) FoldRightColour(zero Colour, fn func(Colour, Thing) Colour) Colour {
	sum := zero
	for i := len(list) - 1; i >= 0; i-- {
		sum = fn(sum, list[i])
	}
	return sum
}

// This methods require Thing be comparable.

// GroupByColour groups elements into a map keyed by Colour.
func (list ThingList) GroupByColour(fn func(Thing) Colour) map[Colour]ThingList {
	result := make(map[Colour]ThingList)
	for _, v := range list {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// These methods require Colour be ordered.

// MinByColour finds the first element which yields the smallest value measured by function fn.
// fn is usually called a projection or measuring function.
// Returns an error if the ThingList is empty.
func (list ThingList) MinByColour(fn func(Thing) Colour) (result Thing, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length ThingList")
		return
	}
	result = list[0]
	if l > 1 {
		min := fn(result)
		for i := 1; i < l; i++ {
			v := list[i]
			f := fn(v)
			if min > f {
				min = f
				result = v
			}
		}
	}
	return
}

// MaxByColour finds the first element which yields the largest value measured by function fn.
// fn is usually called a projection or measuring function.
// Returns an error if the ThingList is empty.
func (list ThingList) MaxByColour(fn func(Thing) Colour) (result Thing, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length ThingList")
		return
	}
	result = list[0]
	if l > 1 {
		max := fn(result)
		for i := 1; i < l; i++ {
			v := list[i]
			f := fn(v)
			if max < f {
				max = f
				result = v
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------
// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.
//-----------------------------------------------------------------------------

// SortWith returns a new ordered ThingList, determined by a func defining ‘less’.
func (list ThingList) SortWith(less func(Thing, Thing) bool) ThingList {
	result := make(ThingList, len(list))
	copy(result, list)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortThingList(result, less, 0, n, maxDepth)
	return result
}

// IsSortedWith reports whether an instance of ThingList is sorted, using the pass func to define ‘less’.
func (list ThingList) IsSortedWith(less func(Thing, Thing) bool) bool {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		if less(list[i], list[i-1]) {
			return false
		}
	}
	return true
}

// SortWithDesc returns a new, descending-ordered ThingList, determined by a func defining ‘less’.
func (list ThingList) SortWithDesc(less func(Thing, Thing) bool) ThingList {
	greater := func(a, b Thing) bool {
		return less(b, a)
	}
	return list.SortWith(greater)
}

// IsSortedDesc reports whether an instance of ThingList is sorted in descending order, using the pass func to define ‘less’.
func (list ThingList) IsSortedWithDesc(less func(Thing, Thing) bool) bool {
	greater := func(a, b Thing) bool {
		return less(b, a)
	}
	return list.IsSortedWith(greater)
}

//-------------------------------------------------------------------------------------------------
// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapThingList(list ThingList, a, b int) {
	list[a], list[b] = list[b], list[a]
}

// Insertion sort
func insertionSortThingList(list ThingList, less func(Thing, Thing) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(list[j], list[j-1]); j-- {
			swapThingList(list, j, j-1)
		}
	}
}

// siftDown implements the heap property on list[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownThingList(list ThingList, less func(Thing, Thing) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(list[first+child], list[first+child+1]) {
			child++
		}
		if !less(list[first+root], list[first+child]) {
			return
		}
		swapThingList(list, first+root, first+child)
		root = child
	}
}

func heapSortThingList(list ThingList, less func(Thing, Thing) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownThingList(list, less, i, hi, first)
	}

	// Pop elements, largest first, into end of list.
	for i := hi - 1; i >= 0; i-- {
		swapThingList(list, first, first+i)
		siftDownThingList(list, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values list[a], list[b], list[c] into list[a].
func medianOfThreeThingList(list ThingList, less func(Thing, Thing) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(list[m1], list[m0]) {
		swapThingList(list, m1, m0)
	}
	if less(list[m2], list[m1]) {
		swapThingList(list, m2, m1)
	}
	if less(list[m1], list[m0]) {
		swapThingList(list, m1, m0)
	}
	// now list[m0] <= list[m1] <= list[m2]
}

func swapRangeThingList(list ThingList, a, b, n int) {
	for i := 0; i < n; i++ {
		swapThingList(list, a+i, b+i)
	}
}

func doPivotThingList(list ThingList, less func(Thing, Thing) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeThingList(list, less, lo, lo+s, lo+2*s)
		medianOfThreeThingList(list, less, m, m-s, m+s)
		medianOfThreeThingList(list, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeThingList(list, less, lo, m, hi-1)

	// Invariants are:
	//	list[lo] = pivot (set up by ChoosePivot)
	//	list[lo <= i < a] = pivot
	//	list[a <= i < b] < pivot
	//	list[b <= i < c] is unexamined
	//	list[c <= i < d] > pivot
	//	list[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(list[b], list[pivot]) { // list[b] < pivot
				b++
			} else if !less(list[pivot], list[b]) { // list[b] = pivot
				swapThingList(list, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(list[pivot], list[c-1]) { // list[c-1] > pivot
				c--
			} else if !less(list[c-1], list[pivot]) { // list[c-1] = pivot
				swapThingList(list, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// list[b] > pivot; list[c-1] < pivot
		swapThingList(list, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeThingList(list, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeThingList(list, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortThingList(list ThingList, less func(Thing, Thing) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortThingList(list, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotThingList(list, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortThingList(list, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortThingList(list, mhi, b)
		} else {
			quickSortThingList(list, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortThingList(list, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortThingList(list, less, a, b)
	}
}
