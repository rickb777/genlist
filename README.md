# GoList

This package is a typewriter for use with Mark Sherman's [gen](https://github.com/clipperhouse/gen), an excellent 
type-driven code generation tool for Go.

This Go code generator offers **List**, **Option** and **Set** containers, loosely influenced bby Scala and Java. 
These have methods for a wide variety of operations such as filtering and sorting, for working with Go types and
slices of types.

The code generator produces flexible type-safe collection classes for your own types. They're not 'generics'
because Go doesn't really have them - but you get the benefit of a clean API for your own classes regardless.

**This work is still beta quality.** If you use it, be prepared for small API changes as it gets completed. 

## Quick start

Start by [installing Go](https://golang.org/dl/), [setting up paths](http://golang.org/doc/code.html), etc. Then:

```
go get github.com/clipperhouse/gen
```

Create a new Go project, and cd into it. Create a main.go file and define a type.

Now, mark it up with a `+gen` annotation in an adjacent comment like so:

```go
// +gen T1 T2
type MyType struct {}
```

where `T1` and `T2` are the templates you want implemented on `MyType`. For this suite, you have a choice of 
`List`, `Option` and `Set`.

Then at the command line, simply type:

```
gen
```

 * If you specified `List`, you should see a new file called mytype_list.go.
 * If you specified `Option`, you should see a new file called mytype_option.go.
 * If you specified `Set`, you should see a new file called mytype_set.go.

Before we examine these in detail, let's look at the unifying interfaces they implement.

## The Collection Interface

The generated code can easily be configured to contain List types and/or Option types and/or Set types. You don't 
*need* them all, but there is are two unifying interfaces provided anyway, and this makes them polymorphic.

A collection looks like this:

```go
// Foo1Collection is an interface for collections of type Foo1, including sets, lists and options (where present).
type Foo1Collection interface {
	// Size gets the size/length of the sequence.
	Size() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo1) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo1))

	// Iter sends all elements along a channel of type Foo1.
	// The first time it is used, order of the elements is not well defined. But the order is stable, which means
	// it will give the same order each subsequent time it is used.
	Iter() <-chan Foo1

	//-------------------------------------------------------------------------
	// Filter returns a new Foo1Collection whose elements return true for a predicate function.
	Filter(predicate func(Foo1) bool) (result Foo1Collection)

	// Partition returns two new Foo1Collections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Foo1) bool) (matching Foo1Collection, others Foo1Collection)

	//-------------------------------------------------------------------------
	// These methods require Foo1 be comparable.

	// Equals verifies that one or more elements of Foo1Collection return true for the passed func.
	Equals(other Foo1Collection) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo1 is not comparable.
	Contains(value Foo1) bool

	//-------------------------------------------------------------------------
	// Sum sums Foo1 elements.
	// Omitted if Foo1 is not numeric.
	Sum() Foo1

	// Mean computes the arithmetic mean of all elements. Panics if the collection is empty.
	// Omitted if Foo1 is not numeric.
	Mean() Foo1

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
```

## The Sequence Interface

List types and Option types are ordered sequences of elements. The principle is that an optional
value is a sequence of zero or one value, whilst a list is a sequence of zero or more.

Sequences are also collections, but not vice versa.

It looks like this:

```go
// Foo1Seq is an interface for sequences of type Foo1, including lists and options (where present).
type Foo1Seq interface {
	Foo1Collection

	// Len gets the size/length of the sequence - an alias for Size()
	Len() int

	//-------------------------------------------------------------------------
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Foo1

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() Foo1

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() Foo1Seq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() Foo1Seq

	//-------------------------------------------------------------------------
	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Foo1) bool) OptionalFoo1

	// Converts the sequence to a list. For lists, this is merely a type assertion.
	ToList() Foo1List

	//-------------------------------------------------------------------------
	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo1 is not comparable.
	Count(value Foo1) int

	// Distinct returns a new Foo1Seq whose elements are all unique.
	// Omitted if Foo1 is not comparable.
	Distinct() Foo1Seq
}
```

### Next: [Lists](List.md)
#### Contents:

 * **Intro**
 * [Lists](List.md)
 * [Options](Option.md)
 * [Joint Lists With Options](Unified.md)

### See also...

* The original [slice](https://clipperhouse.github.io/gen/slice/) in which List was based.
