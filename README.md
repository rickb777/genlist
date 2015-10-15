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

## The List typewriter

The **List** typewriter generates a type and corresponding methods that follow a pattern inspired by Scala's List and
its parent trait [IndexedSeq](http://www.scala-lang.org/api/2.11.7/#scala.collection.IndexedSeq).

Although Go slices are mutable, the generated code *does not provide* any mutation methods, by design. Some functions
make a copy of the list, mutate the copy and return it without changing the original. There is one exception:
A mutating `Swap` method is needed for efficient list sorting; however, the main sort methods operate on copies
of the list.

You can manipulate the underlying slice using normal Go syntax when it needs to be changed, of course.

The **List** typewriter also offers easier ad-hoc sorts.

At it simplest, the annotation looks like:

````go
// +gen List
type Example struct { ... }
````

This creates a core list to hold `Example` values. It provides methods including:

 * IsEmpty, NonEmpty, Len - get simple properties
 * Swap - get a new list with two elements swapped
 * Exists, Forall - tests whether any or all elements match some specified condition
 * Foreach - applies a function to every element in turn, typically causing side-effects
 * Reverse, Shuffle - get a new list that is reversed or shuffled
 * Take, TakeLast, TakeWhile - get a new list without some trailing elements
 * Drop, DropLast, DropWhile - get a new list without some leading elements
 * Filter, Partition - gets a subset, or two disjoint subsets, of the list
 * CountBy, MaxBy, MinBy, DistinctBy - statistics based on supplied operator functions

It adds:

 * Min, Max

but the implementation depends on whether the element type is ordered or not. For ordered elements, Min and Max use
simple inequality operators '<' and '>'. Otherwise a comparison function must be supplied.

Also in the case of ordered elements, the list implements sorting using the [standard Go api](https://golang.org/pkg/sort/)
within these methods:

* Sort, IsSorted
* SortDesc, IsSortedDesc

If the element type is comparable, it adds:

 * Contains, Count - comparison with a specified value
 * Distinct - removal of duplicates

If the element type is numeric, it adds:

 * Sum, Mean

Finally, if a companion option is present, it adds:

 * HeadOption - gets the first element if present

### List For Pointer Elements

A variant is to include a star:

````go
// +gen * List
type Example struct { ... }
````

This creates a basic list of `*Example` pointers, i.e. `[]*Example`.

### Tags

Extra tags can be included to add more features. You can include a comma-separated list of as many tags as you need.

#### `MapTo` Tag

`MapTo[T]` adds code to transform the original list to a new 
list by transforming each element using a function you provide. `MapTo[T]` can be used more than once: 

````go
// +gen List:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

Each tag creates a corresponding `MapToFred`, `MapToJim` etc function.

#### `With[T]` Tag

`With[T]` adds methods that need functions depending on another type, `T`.

````go
// +gen List:"With[Colour], With[Texture]"
type Example struct { ... }
````

This adds methods according to the type `T`:

 * FoldLeftT, FoldRightT (always) - provide general summation functions
 * GroupByT (only if `T` is comparable)
 * SumT, MeanT (only if `T` is numeric)
 * MinT, MaxT (only if `T` is ordered)

#### `SortWith` Tag

This adds extra sort methods that simply depend on providing a comparator function.

````go
// +gen List:"SortWith"
type Example struct { ... }
````


### See also...

* The original [slice](https://clipperhouse.github.io/gen/slice/) in which List was based.
