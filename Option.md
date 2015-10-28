# GoList - Options

## The Option typewriter

The **Option** typewriter generates a type and corresponding methods that follow a pattern inspired by 
Scala's [Option](http://www.scala-lang.org/api/2.11.7/#scala.Option).

By design, the generated code *does not provide* any mutation methods.

Options can store a value or a pointer to a value. So if you choose to store a pointer value, using this pointer,
you can change the value itself should you need to.

At it simplest, the annotation looks like:

````go
// +gen Option
type Example struct { ... }
````

#### Core methods

This creates a optional example type `OptionalExample` to hold `Example` values. It provides all the
*T*Collection methods, including:

 * **Get**, **GetOrElse**, **OrElse** - get the current value
 * **Head** - an alias for **Get**
 * **Len**, **Size** - get simple properties (these are aliases)
 * **IsEmpty**, **NonEmpty** - get simple properties
 * **IsDefined** - an alias for NonEmpty
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the option
 * **Send** - get a channel that supplies values in sequence
 * **MkString**, **MkString3**, **String** - constructs a string representation of the option (just like a list)

#### Comparable Methods

If the element type is *comparable*, it adds:

 * **Equals** - compare with another list
 * **Contains**, **Count** - compare with a specified value
 * **Distinct** - remove duplicates

### Option For Pointer Elements

A variant is to include a star:

````go
// +gen * Option
type Example struct { ... }
````

This creates a methods that access the contained element as a `*Example` pointer instead of by value.

### Tags

Extra tags can be included to add more features. You can include a comma-separated list of as many tags as you need.

#### `List` and `Set` Tags

`List` adds list implementation on the same type (e.g. `Example`). Some additional methods are provided to convert
between options and lists.

`Set` adds set implementation on the same type (e.g. `Example`).  Some additional methods are provided to convert
between options and sets.

#### `MapTo` Tag

`MapTo[T]` adds code to transform the original option to a new 
option by transforming its element using a function you provide. `MapTo[T]` can be used more than once: 

````go
// +gen Option:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

Each tag creates a corresponding `MapToFred`, `MapToJim` etc function. These functions return the corresponding
`OptionalFred`, `OptionalJim` etc values.

In addition, a corresponding `FlatMapToFred`, `FlatMapToJim` etc function is generated. These functions return the 
corresponding `OptionalFred`, `OptionalJim` etc values. A mapping function transforms individual `Example` item to a
sequence of its `Fred` equivalent. This sequence is added to the result set.

### Next: [Sets](Set.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * **Options**
 * [Sets](Set.md)
 * [Joint Lists with Options  and/or Sets](Unified.md)
 * [Plumbing functions](Plumbing.md)
