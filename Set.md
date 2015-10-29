# GoList - Sets

Sets hold hold a group of items without duplicates. There is no obvious order of the elements in the set.

Normally, the generated code *does not provide* any mutation methods (but this can be modified, as described below).

Set methods include:

 * **IsEmpty**, **NonEmpty**, **Size** - get simple properties
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Equals** - compare with another set
 * **Union**, **Intersection**, **Difference** - provides standard set algebra
 * **Contains** - compare with a specified value
 * **Add**, **Remove** - creates a modified set (n.b. these might mutate the set - see below)
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the set
 * **Iter** - get a channel that supplies values in sequence
 * **CountBy**, **MaxBy**, **MinBy**, **DistinctBy** - statistics based on supplied operator functions
 * **MkString**, **MkString3**, **String** - constructs a string representation of the set (like a list)

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

### Next: [Joint Lists with Options and/or Sets](Unified.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * **Sets**
 * [Joint Lists with Options and/or Sets](Unified.md)
 * [Plumbing functions](Plumbing.md)
