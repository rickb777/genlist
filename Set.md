# GenList - Sets

Sets hold hold a group of items without duplicates. There is no obvious order of the elements in the set.

Normally, the generated code *does not provide* any mutation methods.

Tags can, of course, be included in the quotes as a comma-separated list as per usual. There is only one
tag implemented at present:

 * Mutate: changes the Add and Remove methods to mutate the current set, instead of returning a new set.

Set methods include:

 * **IsEmpty**, **NonEmpty**, **Size** - get simple properties
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Equals** - compare with another set
 * **Union**, **Intersection**, **Difference** - provides standard set algebra
 * **Contains** - compare with a specified value
 * **Add**, **Remove** - creates a modified set
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the set
 * **Iter** - get a channel that supplies values in sequence
 * **CountBy**, **MaxBy**, **MinBy**, **DistinctBy** - statistics based on supplied operator functions
 * **MkString**, **MkString3**, **String** - constructs a string representation of the set (like a list)

### Tags

Extra tags can be included to add more features. You can include a comma-separated list of as many tags as you need.

#### `List` Tag

`List` adds list implementation on the same type (e.g. `Example`).

````go
// +gen Set:"List"
type Example struct { ... }
````

Some additional methods are provided to convert between sets and lists. Therefore, it is better to use this tag
rather than generate separate data structures using `+gen Set List` or similar.

#### `MapTo` Tag

`MapTo[T]` adds code to transform the original set to a new 
set by transforming its element using a function you provide. `MapTo[T]` can be used more than once: 

````go
// +gen Set:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

Each tag creates a corresponding `MapToFred`, `MapToJim` etc function. These functions return the corresponding
`FredSet`, `JimSet` etc values. A mapping function transforms individual `Example` item to its `Fred` equivalent.

In addition, a corresponding `FlatMapToFred`, `FlatMapToJim` etc function is generated. These functions return the 
corresponding `FredSet`, `JimSet` etc values. A mapping function transforms individual `Example` item to a
sequence of its `Fred` equivalent. This sequence is added to the result set.

### Next: [Joint Lists With Options](Unified.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * **Sets**
 * [Joint Lists With Options](Unified.md)
