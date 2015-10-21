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

### Next: [Joint Lists With Options](Unified.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * **Sets**
 * [Joint Lists With Options](Unified.md)
