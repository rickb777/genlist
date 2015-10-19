# GenList - Options

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
*T*Seq methods, including:

 * **Head**, **Tail** - get the first element and the rest
 * **Init**, **Last** - get the last element and the rest
 * **IsEmpty**, **NonEmpty**, **Len** - get simple properties
 * **IsDefined** - an alias for NonEmpty
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Reverse**, **Shuffle** - get a new list that is reversed or shuffled
 * **Take**, **TakeWhile**, **DropLast** - get a new list without some trailing elements
 * **Drop**, **DropWhile**, **TakeLast** - get a new list without some leading elements
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the list

#### Comparable Methods

If the element type is *comparable*, it adds:

 * **Equals** - compare with another list
 * **IndexOf**, **IndexOf2** - find the index of the first match
 * **LastIndexOf**, **LastIndexOf2** - find the index of the last match
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

#### `MapTo` Tag

`MapTo[T]` adds code to transform the original option to a new 
option by transforming its element using a function you provide. `MapTo[T]` can be used more than once: 

````go
// +gen Option:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

Each tag creates a corresponding `MapToFred`, `MapToJim` etc function. These functions return the corresponding
`OptionalFred`, `OptionalJim` etc values.

### Next: [Joint Lists With Options](Unified.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * **Options**
 * [Using Lists and Options Together](Unified.md)
