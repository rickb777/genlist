# GenList - Using Lists and Options Together

List hold sequences of zero or more items. Options hold sequences of zero or one item. Although the similarity between
them should not be overstated, they still provide useful generalities. This is why the [sequence interface](README.md)
is provided.

When a type has both its list and its option generated, extra methods are provided and these are described below.
Obtaining both is easy, simply include one as the required type and the other as a tag, e.g.

````go
// +gen Option:"List"
type Example struct { ... }
````
or, equivalently,

````go
// +gen List:"Option"
type Example struct { ... }
````

Other tags can, of course, be included in the quotes as a comma-separated list as per usual.

This creates 

 1. an optional example type `OptionalExample`, and 
 2. the list example type `ExampleList`, plus 
 3. the [common interface](README.md) `ExampleSeq`
 
The common interface includes:

 * **Head**, **IsEmpty**, **NonEmpty**, **Len** - get simple properties
 * **Exists**, **Forall** - tests whether any or all elements match some specified condition
 * **Foreach** - applies a function to every element in turn, typically causing side-effects
 * **Filter**, **Partition** - gets a subset, or two disjoint subsets, of the list
 * **Contains**, **Count** - comparison with a specified value (only if *comparable*)
 * **Distinct** - removes duplicates (only if *comparable*)
 * **Find** - finds a match in the sequence, if there is one.
 * **ToList** - returns the sequence as a list (for lists, this is a type conversion only because it's a list already)

In addition, the listDocumentatio has the following extra methods:

 * **HeadOption**, **TailOption** - gets the first/last item in the list, if there is one.

#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * **Joint Lists With Options**
