# GoList - Using Lists and Options and Sets Together

The code generator can be used for combinations of lists, options and sets.

List hold sequences of zero or more items. Options hold sequences of zero or one item. Although the similarity between
them should not be overstated, they still provide useful generalities. This is why the [sequence interface](README.md)
is provided.

Sets provide collections that have no order and in which there are no duplicates.

When a type has a combination of list, option and set generated, extra methods are provided and these are described
below. Obtaining these combinations is easy, simply include one as the required type and the other as a tag, e.g.

````go
// +gen Option:"List Set"
type Example struct { ... }
````
or, equivalently,

````go
// +gen List:"Option Set"
type Example struct { ... }
````
or, equivalently,

````go
// +gen Set:"List Option"
type Example struct { ... }
````

Other tags can, of course, be included in the quotes as per usual.

This creates 

 1. an optional example type `OptionalExample`, and 
 2. the list example type `ExampleList`, plus 
 3. the set example type `ExampleSet`, plus 
 4. the [common interfaces](README.md) `ExampleCollection` and `ExampleSeq`

The list has the following extra methods:

 * **HeadOption**, **TailOption** - gets the first/last item in the list, if there is one.
 * **Find** - finds an item in the list, if there is one.
 * **ToSet** - returns a set containing the list elements, without duplicates.

The option has the following extra methods:

 * **ToList** - returns a list containing the option element, if there is one.

The set has the following extra methods:

 * **ToList** - returns a list containing the set elements in some arbitrary order.

### Next: [Plumbing functions](Plumbing.md)
#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * [Sets](Set.md)
 * **Joint Lists with Options and/or Sets**
 * [Plumbing functions](Plumbing.md)
