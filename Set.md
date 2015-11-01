# GoList - Sets

Sets hold hold a group of items without duplicates. There is no obvious order of the elements in the set.

Normally, the generated code *does not provide* any mutation methods (but this can be modified, as described below).

At it simplest, the annotation looks like:

````go
// +gen Set
type Example struct { ... }
````

#### Core methods

The generated code is a set to hold `Example` values. It provides methods including:

 * **IsEmpty**, **NonEmpty** - get simple properties
 * **Size**, **Len** - get simple properties (these are aliases)
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the collection
 * **CountBy**, **MaxBy**, **MinBy**, **DistinctBy** - statistics based on supplied operator functions
 * **MkString**, **MkString3**, **String** - constructs a string representation of the collection
 * **Send** - get a channel that supplies values in sequence

#### Set Methods

 * **Union**, **Intersection**, **Difference** - provides standard set algebra
 * **Contains** - compare with a specified value
 * **Add**, **Remove** - creates a modified set (n.b. these might mutate the set - see below)

#### Comparable Methods

If the element type is *comparable*, it adds:

 * **Equals** - compare with another collection
 * **Contains** - compare with a specified value

#### Numeric Methods

If the element type is *numeric*, it adds:

 * **Sum** - sum all elements
 * **Mean** - compute the arithmetic mean

#### Min and Max

It always adds:

 * **Min**, **Max** - find the minimum/maximum value

but the implementation depends on whether the element type is *ordered* or not. For ordered elements, Min and Max use
simple inequality operators '<' and '>'. Otherwise a comparison function must be supplied.

### Tags

Extra tags can be included to add more features. You can include a space-separated list of as many tags as you need.

#### `Mutate` Tag

The Mutate changes the Add and Remove methods to mutate the current set, instead of returning a new set. This
might be used as a performance optimisation when dealing with large sets.

#### `List` Tag

`List` adds list implementation on the same type (e.g. `Example`).

````go
// +gen Set:"List"
type Example struct { ... }
````

Some additional methods are provided to convert between sets and lists. Therefore, it is better to use this tag
rather than generate separate data structures using `+gen Set List` or similar.

#### `MapTo` Tag

`MapTo[T]` adds code to transform the original set to a new set by transforming its element using a function you provide.
`MapTo[T]` can be used more than once: 

````go
// +gen Set:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

This adds methods according to the parameter type `T` (e.g. `Fred` and `Jim` above):

 * **MapTo**T - converts to type `T`
 * **FlatMapTo**T - converts to collections of `T`

For `MapToT`, a mapping function transforms an individual `Example` item to its `Fred` equivalent.

For `FlatMapToT`, a mapping function transforms an individual `Example` item to a
sequence of its `T` equivalent. This sequence is added to the result set.

Normally the result is a `TCollection`. 
But if the target type is a basic type, such as `string` or `int`, the result is a slice of the basic type instead.

#### `With[T]` Tag

`With[T]` adds methods that need functions depending on another type, `T`.

````go
// +gen Set:"With[Colour], With[Texture]"
type Example struct { ... }
````

This adds methods according to the type `T`:

 * **FoldLeft**T, **FoldRight**T (always) - provide general summation functions
 * **GroupBy**T (only if `T` is *comparable*) - clustering and histogramming
 * **Sum**T, **Mean**T (only if `T` is *numeric*)
 * **MinBy**T, **MaxBy**T (only if `T` is *ordered*) - find the min/max list entry according to a projection function

### Next: [Options](Option.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * **Sets**
 * [Options](Option.md)
 * [Joint Lists with Options and/or Sets](Unified.md)
 * [Plumbing functions](Plumbing.md)
