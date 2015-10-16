# GenList

This package is a typewriter for use with Mark Sherman's [gen](https://github.com/clipperhouse/gen), an excellent 
type-driven code generation tool for Go.

This Go code generator offers scala-like **List** and **Option** containers, with methods such as filtering and sorting,
for working with Go types and slices of types.

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
// FooSeq is an interface for sequences of type Foo, including lists and options (where present).
type FooSeq interface {
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Foo

	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo))

	// Filter returns a new FooSeq whose elements return true for a predicate function.
	Filter(predicate func(Foo) bool) (result FooSeq)

	// Partition returns two new FooLists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Foo) bool) (matching FooSeq, others FooSeq)

	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Foo) bool) OptionalFoo

	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Foo is not comparable.
	Equals(other FooSeq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo is not comparable.
	Contains(value Foo) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo is not comparable.
	Count(value Foo) int
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
