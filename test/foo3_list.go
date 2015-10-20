// Generated by: setup
// TypeWriter: List
// Directive: +test on *Foo3

package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
)

// Foo3Seq is an interface for sequences of type *Foo3, including lists and options (where present).
type Foo3Seq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	//-------------------------------------------------------------------------
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() *Foo3

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() *Foo3

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() Foo3Seq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() Foo3Seq

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(*Foo3) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(*Foo3) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(*Foo3))

	//-------------------------------------------------------------------------
	// Filter returns a new Foo3Seq whose elements return true for a predicate function.
	Filter(predicate func(*Foo3) bool) (result Foo3Seq)

	// Partition returns two new Foo3Lists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(*Foo3) bool) (matching Foo3Seq, others Foo3Seq)

	//-------------------------------------------------------------------------
	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(*Foo3) bool) OptionalFoo3

	// Converts the sequence to a list. For lists, this is merely a type assertion.
	ToList() Foo3List

	//-------------------------------------------------------------------------
	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Foo3 is not comparable.
	Equals(other Foo3Seq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo3 is not comparable.
	Contains(value *Foo3) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo3 is not comparable.
	Count(value *Foo3) int

	// Distinct returns a new Foo3Seq whose elements are all unique.
	// Omitted if Foo3 is not comparable.
	Distinct() Foo3Seq
}

//-------------------------------------------------------------------------------------------------
// Foo3List is a slice of type *Foo3. Use it where you would use []*Foo3.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Foo3List []*Foo3

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// panics if list is empty
func (list Foo3List) Head() *Foo3 {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// panics if list is empty
func (list Foo3List) Last() *Foo3 {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// panics if list is empty
func (list Foo3List) Tail() Foo3Seq {
	return Foo3List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// panics if list is empty
func (list Foo3List) Init() Foo3Seq {
	return Foo3List(list[:len(list)-1])
}

// IsEmpty tests whether Foo3List is empty.
func (list Foo3List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Foo3List is empty.
func (list Foo3List) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list Foo3List) ToList() Foo3List {
	return list
}

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list Foo3List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list Foo3List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// Exists verifies that one or more elements of Foo3List return true for the passed func.
func (list Foo3List) Exists(fn func(*Foo3) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Foo3List return true for the passed func.
func (list Foo3List) Forall(fn func(*Foo3) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Foo3List and executes the passed func against each element.
func (list Foo3List) Foreach(fn func(*Foo3)) {
	for _, v := range list {
		fn(v)
	}
}

// Reverse returns a copy of Foo3List with all elements in the reverse order.
func (list Foo3List) Reverse() Foo3List {
	numItems := len(list)
	result := make(Foo3List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of Foo3List, using a version of the Fisher-Yates shuffle.
func (list Foo3List) Shuffle() Foo3List {
	numItems := len(list)
	result := make(Foo3List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

// Take returns a new Foo3List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo3List) Take(n int) Foo3List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new Foo3List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo3List) Drop(n int) Foo3List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new Foo3List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo3List) TakeLast(n int) Foo3List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new Foo3List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo3List) DropLast(n int) Foo3List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new Foo3List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list Foo3List) TakeWhile(p func(*Foo3) bool) (result Foo3List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new Foo3List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list Foo3List) DropWhile(p func(*Foo3) bool) (result Foo3List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Filter returns a new Foo3List whose elements return true for func.
func (list Foo3List) Filter(fn func(*Foo3) bool) Foo3Seq {
	result := make(Foo3List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new Foo3Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list Foo3List) Partition(p func(*Foo3) bool) (Foo3Seq, Foo3Seq) {
	matching := make(Foo3List, 0, len(list)/2)
	others := make(Foo3List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of Foo3List that return true for the passed predicate.
func (list Foo3List) CountBy(predicate func(*Foo3) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Foo3List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list Foo3List) MinBy(less func(*Foo3, *Foo3) bool) (result *Foo3, err error) {
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

// MaxBy returns an element of Foo3List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list Foo3List) MaxBy(less func(*Foo3, *Foo3) bool) (result *Foo3, err error) {
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

// DistinctBy returns a new Foo3List whose elements are unique, where equality is defined by a passed func.
func (list Foo3List) DistinctBy(equal func(*Foo3, *Foo3) bool) (result Foo3List) {
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
func (list Foo3List) IndexWhere(p func(*Foo3) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list Foo3List) IndexWhere2(p func(*Foo3) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list Foo3List) LastIndexWhere(p func(*Foo3) bool) int {
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
func (list Foo3List) LastIndexWhere2(p func(*Foo3) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

// These methods require *Foo3 be comparable.

// Equals verifies that one or more elements of Foo3List return true for the passed func.
func (list Foo3List) Equals(other Foo3Seq) bool {
	if len(list) != other.Len() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a *Foo3) {
		if eq {
			v := list[i]
			if *v != *a {
				eq = false
			}
			i += 1
		}
	})
	return eq
}

// These methods require *Foo3 be comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list Foo3List) IndexOf(value *Foo3) int {
	for i, v := range list {
		if *v == *value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list Foo3List) IndexOf2(value *Foo3, from int) int {
	for i, v := range list {
		if i >= from && *v == *value {
			return i
		}
	}
	return -1
}

// Contains verifies that a given value is contained in Foo3List.
func (list Foo3List) Contains(value *Foo3) bool {
	for _, v := range list {
		if *v == *value {
			return true
		}
	}
	return false
}

// Count gives the number elements of Foo3List that match a certain value.
func (list Foo3List) Count(value *Foo3) (result int) {
	for _, v := range list {
		if *v == *value {
			result++
		}
	}
	return
}

// Distinct returns a new Foo3List whose elements are unique.
func (list Foo3List) Distinct() Foo3Seq {
	result := make(Foo3List, 0)
	appended := make(map[Foo3]bool)
	for _, v := range list {
		if !appended[*v] {
			result = append(result, v)
			appended[*v] = true
		}
	}
	return result
}

// Min returns the first element of Foo3List containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Returns an error if the Foo3List is empty.
func (list Foo3List) Min(less func(*Foo3, *Foo3) bool) (result *Foo3, err error) {
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

// Max returns the first element of Foo3List containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Returns an error if the Foo3List is empty.
func (list Foo3List) Max(less func(*Foo3, *Foo3) bool) (result *Foo3, err error) {
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

// String implements the Stringer interface to render the list as a comma-separated array.
func (list Foo3List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (list Foo3List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (list Foo3List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", *v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", *v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}

// optionForList

// First returns the first element that returns true for the passed func. Returns error if no elements return true.
func (list Foo3List) Find(fn func(*Foo3) bool) OptionalFoo3 {
	for _, v := range list {
		if fn(v) {
			//return SomeFoo3(v)
		}
	}
	return NoFoo3()
}

// HeadOption gets the first item in the list, provided there is one.
func (list Foo3List) HeadOption() OptionalFoo3 {
	l := len(list)
	if l > 0 {
		return SomeFoo3(list[0])
	} else {
		return NoFoo3()
	}
}

// TailOption gets the last item in the list, provided there is one.
func (list Foo3List) LastOption() OptionalFoo3 {
	l := len(list)
	if l > 0 {
		return SomeFoo3(list[l-1])
	} else {
		return NoFoo3()
	}
}

//-------------------------------------------------------------------------------------------------
// OptionalFoo3 is an optional of type *Foo3. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalFoo3 struct {
	x *Foo3
}

// shared none value
var noneFoo3 = OptionalFoo3{nil}

func NoFoo3() OptionalFoo3 {
	return noneFoo3
}

func SomeFoo3(x *Foo3) OptionalFoo3 {

	if x == nil {
		return NoFoo3()
	}
	return OptionalFoo3{x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalFoo3) Head() *Foo3 {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return (o.x)
}

// panics if option is empty
func (o OptionalFoo3) Last() *Foo3 {
	return o.Head()
}

// panics if option is empty
func (o OptionalFoo3) Tail() Foo3Seq {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return noneFoo3
}

// panics if option is empty
func (o OptionalFoo3) Init() Foo3Seq {
	return o.Tail()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo3) Get() *Foo3 {
	return o.Head()
}

func (o OptionalFoo3) GetOrElse(d func() *Foo3) *Foo3 {
	if o.IsEmpty() {
		return d()
	}
	return o.x
}

func (o OptionalFoo3) OrElse(alternative func() OptionalFoo3) OptionalFoo3 {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo3) Len() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalFoo3) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalFoo3) NonEmpty() bool {
	return o.x != nil
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalFoo3) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo3) Find(predicate func(*Foo3) bool) OptionalFoo3 {
	if o.IsEmpty() {
		return o
	}
	if predicate(o.x) {
		return o
	}
	return noneFoo3
}

func (o OptionalFoo3) Exists(predicate func(*Foo3) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(o.x)
}

func (o OptionalFoo3) Forall(predicate func(*Foo3) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(o.x)
}

func (o OptionalFoo3) Foreach(fn func(*Foo3)) {
	if o.NonEmpty() {
		fn(o.x)
	}
}

func (o OptionalFoo3) Filter(predicate func(*Foo3) bool) Foo3Seq {
	return o.Find(predicate)
}

func (o OptionalFoo3) Partition(predicate func(*Foo3) bool) (Foo3Seq, Foo3Seq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(o.x) {
		return o, noneFoo3
	}
	return noneFoo3, o
}

//-------------------------------------------------------------------------------------------------
// These methods require *Foo3 be comparable.

// Equals verifies that one or more elements of Foo3List return true for the passed func.
func (o OptionalFoo3) Equals(other Foo3Seq) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Len() > 1 {
		return false
	}
	a := o.Head()
	b := other.Head()
	return *a == *b
}

func (o OptionalFoo3) Contains(value *Foo3) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == *value
}

func (o OptionalFoo3) Count(value *Foo3) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new Foo3Seq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if Foo3 is not comparable.
func (o OptionalFoo3) Distinct() Foo3Seq {
	return o
}

func (o OptionalFoo3) ToList() Foo3List {
	if o.IsEmpty() {
		return Foo3List{}
	}
	return Foo3List{o.x}
}

//-------------------------------------------------------------------------------------------------
// String implements the Stringer interface to render the option as an array of one element.
func (o OptionalFoo3) String() string {
	return o.MkString(",")
}

// MkString concatenates the values as a string.
func (o OptionalFoo3) MkString(sep string) string {
	return o.MkString3("[", sep, "]")
}

// MkString3 concatenates the values as a string.
func (o OptionalFoo3) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}
