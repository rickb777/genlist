# GoList

This package is a typewriter for use with Mark Sherman's [gen](https://github.com/clipperhouse/gen), an excellent 
type-driven code generation tool for Go.

This Go code generator offers scala-like **List** and **Option** containers, with methods such as filtering and sorting,
for working with Go types and slices of types.

It produces flexible type-safe collection classes for your own types. They're not 'generics' because Go doesn't have
them - but you get the benefit of a clean API for your own classes regardless.

**This work is still beta quality.** If you use it, be prepared for small API changes as it gets completed. 

## Quick start

Start by [installing Go](https://golang.org/dl/), [setting up paths](http://golang.org/doc/code.html), etc. Then:

```
go get github.com/clipperhouse/gen
```

Create a new Go project, and cd into it. Create a main.go file and define a type.

Now, mark it up with a +gen annotation in an adjacent comment like so:

```go
// +gen List
type MyType struct {}
```

And at the command line, simply type:

```
gen
```

You should see a new file, named mytype_list.go. Have a look around.

## The Sequence Interface

The generated code can easily be configured to contain List types and/or Option types. You don't *need* both, but
there is a unifying interface provided anyway, and this makes them polymorphic. The principle is that an optional
value is a sequence of zero or one value, whilst a list is a sequence of zero or more. 

It looks like this:

```go
// Foo1Seq is an interface for sequences of type Foo1, including lists and options (where present).
type Foo1Seq interface {
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Foo1

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() Foo1

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() Foo1Seq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() Foo1Seq

	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo1) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo1))

	// Filter returns a new Foo1Seq whose elements return true for a predicate function.
	Filter(predicate func(Foo1) bool) (result Foo1Seq)

	// Partition returns two new Foo1Lists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Foo1) bool) (matching Foo1Seq, others Foo1Seq)

	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Foo1) bool) OptionalFoo1

	// Converts the sequence to a list. For lists, this is merely a type assertion.
	ToList() Foo1List

	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Foo1 is not comparable.
	Equals(other Foo1Seq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo1 is not comparable.
	Contains(value Foo1) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo1 is not comparable.
	Count(value Foo1) int

	// Distinct returns a new Foo1Seq whose elements are all unique.
	// Omitted if Foo1 is not comparable.
	Distinct() Foo1Seq

	// Sum sums Foo1 elements.
	// Omitted if Foo1 is not numeric.
	Sum() Foo1
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
