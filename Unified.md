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

 * an optional example type `OptionalExample`, and 
 * the list example type `ExampleList`, plus 
 * the [common interface](README.md) `ExampleSeq`.

In addition, `ExampleList` has the following extra methods:

 * **HeadOption** - gets the first item in the list, if there is one.
 * **Find** - finds a match in the list, if there is one.
 * **ToList** - returns the list (a type conversion only because it's a list already)

In addition, `OptionalExample` has the following extra methods:

 * **Find** - finds a match in the list, if there is one.
 * **ToList** - returns the option as a list

#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * **Joint Lists With Options**
