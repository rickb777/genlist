# GenList

This package is a typewriter for use with [gen](https://github.com/clipperhouse/gen), a type-driven code generation tool for Go.

This Go code generator offers scala-like **List** and **Option** containers, with methods such as filtering and sorting,
for working with Go types and slices of types.

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
 * Foreach
 * Reverse, Shuffle - get a new list that is reversed or shuffled
 * Take, TakeLast, TakeWhile - get a new list without some trailing elements
 * Drop, DropLast, DropWhile - get a new list without some leading elements
 * Filter, Partition
 * CountBy, MaxBy, MinBy, DistinctBy

If the element type is comparable, it adds:

 * Contains, Count, Distinct

If the element type is numeric, it adds:

 * Sum, Mean

It adds:

 * Min, Max

but the implementation depends on whether the element type is ordered or not. For ordered elements, Min and Max use
simple inequality operators '<' and '>'. Otherwise a comparison function must be supplied.

Also in the case of ordered elements, it implements sorting using the standard Go api in these methods:

* Sort, IsSorted
* SortDesc, IsSortedDesc

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
