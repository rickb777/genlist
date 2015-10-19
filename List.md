# GenList - Lists

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

#### Core methods

The generated code is a list to hold `Example` values. It provides methods including:

 * **Head**, **Tail** - get the first element and the rest
 * **Init**, **Last** - get the last element and the rest
 * **IsEmpty**, **NonEmpty**, **Len** - get simple properties
 * **Swap** - get a new list with two elements swapped
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Reverse**, **Shuffle** - get a new list that is reversed or shuffled
 * **Take**, **TakeWhile**, **DropLast** - get a new list without some trailing elements
 * **Drop**, **DropWhile**, **TakeLast** - get a new list without some leading elements
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the list
 * **CountBy**, **MaxBy**, **MinBy**, **DistinctBy** - statistics based on supplied operator functions
 * **IndexWhere**, **IndexWhere2** - find the index of the first match using a predicate function
 * **LastIndexWhere**, **LastIndexWhere2** - find the index of the last match using a predicate function

#### Comparable Methods

If the element type is *comparable*, it adds:

 * **Equals** - compare with another list
 * **IndexOf**, **IndexOf2** - find the index of the first match
 * **LastIndexOf**, **LastIndexOf2** - find the index of the last match
 * **Contains**, **Count** - compare with a specified value
 * **Distinct** - remove duplicates

#### Numeric Methods

If the element type is *numeric*, it adds:

 * **Sum** - sum all elements
 * **Sum** - compute the arithmetic mean

#### Min and Max

It always adds:

 * **Min**, **Max** - find the minimum/maximum value

but the implementation depends on whether the element type is ordered or not. For ordered elements, Min and Max use
simple inequality operators '<' and '>'. Otherwise a comparison function must be supplied.

#### Sorting

Also in the case of *ordered* elements, the list implements sorting using the [standard Go api](https://golang.org/pkg/sort/)
within these methods:

* **Sort**, **IsSorted** - sort in ascending order
* **SortDesc**, **IsSortedDesc** - sort in descending order

#### Optional Extremae

Finally, if a companion option is present, it adds:

 * **HeadOption** - get the first element if present (similar to Head)
 * **LastOption** - get the last element if present (similar to Last)

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

 * **FoldLeft**T, **FoldRight**T (always) - provide general summation functions
 * **GroupBy**T (only if `T` is *comparable*) - clustering and histogramming
 * **Sum**T, **Mean**T (only if `T` is *numeric*)
 * **MinBy**T, **MaxBy**T (only if `T` is *ordered*) - find the min/max list entry according to a projection function

#### `SortWith` Tag

This adds extra sort methods (as alternatives) that simply depend on providing a comparator function. The benefit is
simpler usage at the calling sites, but much more code is generated to implement the sorting, so there is a trade-off
you can decide upon.

````go
// +gen List:"SortWith"
type Example struct { ... }
````

 * **SortWith**, **IsSortedWith**
 * **SortWithDesc**, **IsSortedWithDesc**

### Next: [Options](Option.md)
#### Contents:

 * [Intro](README.md)
 * **Lists**
 * [Options](Option.md)
 * [Joint Lists With Options](Unified.md)
