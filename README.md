# GenList

This package is a typewriter for use with [gen](https://github.com/clipperhouse/gen), a type-driven code generation tool for Go.

This Go code generator offers scala-like **List** and **Option** containers, with methods such as filtering and sorting,
for working with Go types and slices of types.

## The Sequence Interface

The generated code can easily be configured to contain List types and/or Option types. You don't *need* both, but
there is a unifying interface provided anyway, and this makes them polymorphic. The principle is that an optional
value is a sequence of zero or one value, whilst a list is a sequence of zero or more. 

It looks like this:

```go
// ExampleSeq is an interface for sequences of type Example, including lists and options (where present).
type ExampleSeq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predictate supplied.
	Exists(predicate func(Example) bool) bool

	// Forall returns true if every element in the sequence matches the predictate supplied.
	Forall(predicate func(Example) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Example))

	// Filter returns a new ExampleSeq whose elements return true for func.
	Filter(predicate func(Example) bool) (result ExampleSeq)

	// Converts the sequence to a list. For lists, this is a no-op.
	ToList() ExampleList

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Example is not comparable.
	Contains(value Example) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Example is not comparable.
	Count(value Example) int
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
