# GoList

This package is a typewriter for use with Mark Sherman's [gen](https://github.com/clipperhouse/gen), an excellent 
type-driven code generation tool for Go.

This Go code generator offers **List**, **Option** and **Set** containers, loosely influenced by Scala and Java. 
These have methods for a wide variety of operations such as filtering and sorting, for working with Go types and
slices of types.

A fourth module provides typesafe standard **plumbing** functions for channels that carry streams of particular types.

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
 * If you specified `Plumbing`, you should see a new file called mytype_plumbing.go.

Before we examine these in detail, let's look at the unifying interfaces they implement.

## The Collection Interface

The generated code can easily be configured to contain List types and/or Option types and/or Set types. You don't 
*need* them all, but there is is a unifying interface provided anyway, and this makes them polymorphic.

All three (lists, options and sets) are collections. A collection looks like this, for a type `Num1`:

```go
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

	//-------------------------------------------------------------------------
	// ToSlice returns a plain slice containing all the elements in the collection.
	// This is useful for bespoke iteration etc.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	ToSlice() []Num1

	// ToList gets all the elements in a in List.
	ToList() Num1List

	// Send sends all elements along a channel of type Num1.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	Send() <-chan Num1

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(Num1) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(Num1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Num1))

	//-------------------------------------------------------------------------
	// Filter returns a new Num1Collection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(Num1) bool) (result Num1Collection)

	// Partition returns two new Num1Collections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Num1) bool) (matching Num1Collection, others Num1Collection)

	//-------------------------------------------------------------------------

	// Equals verifies that another Num1Collection has the same size and elements as this one. Also,
	// if the collection is a sequence, the order must be the same.
	// Omitted if Num1 is not comparable.
	Equals(other Num1Collection) bool

	// Contains tests whether a given value is present in the collection.
	// Omitted if Num1 is not comparable.
	Contains(value Num1) bool

	//-------------------------------------------------------------------------
	// Sum sums Num1 elements.
	// Omitted if Num1 is not numeric.
	Sum() Num1

	// Mean computes the arithmetic mean of all elements. Panics if the collection is empty.
	// Omitted if Num1 is not numeric.
	Mean() Num1

	// Min returns the minimum value of Num1List. In the case of multiple items being equally minimal,
	// the first such element is returned. Panics if the collection is empty.
	Min() Num1

	// Max returns the maximum value of Num1List. In the case of multiple items being equally maximal,
	// the first such element is returned. Panics if the collection is empty.
	Max() Num1

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

There are slight variants in the generated code. This is dependent on the element type: 

 * For comparable types, methods suitable for comparisons are included. 
 * For ordered types, methods using greater than/less than operations are included. 
 * For numeric types, summation and average methods are included. 

### Next: [Lists](List.md)
#### Contents:

 * **Intro**
 * [Lists](List.md)
 * [Options](Option.md)
 * [Joint Lists with Options and/or Sets](Unified.md)
 * [Plumbing functions](Plumbing.md)

### See also...

* The original [slice](https://clipperhouse.github.io/gen/slice/) in which List was based.
