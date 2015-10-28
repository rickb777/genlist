# GoList - Plumbing

Channel communications often provide simple decoupling and obvious concurrency. 
This plumbing suite provides standard functions for piping data between goroutines.
It consists of a suite of standard transformation functions, code-generated in a typesafe way, intended to be
plumbed into app-specific sources and sinks for the streams of data.

All these functions run until the input channel is closed. They then close their output channel(s).

None of these functions create a goroutine - this must be done at the call site.

For a type `T`, methods include:

 * **TDelta** - duplicate a stream to two channels.
 * **TZip** - interleaves two streams onto a single output.
 * **TBlackhole** - silently consume a stream.
 * **TFilter** - passes along only those elements of the stream that match a given predicate.
 * **TPartition** - divides the stream into two channels, one for elements that matches a given predicate, and the
   other for those that don't.
 * **TMap** - transforms a stream by applying a function to each element, where the function maps one input element
   to one output element.
 * **TFlatMap** - transforms a stream by applying a function to each element, where the function maps one input
   element to zero or more output elements.

### Tags

Extra tags can be included to add more features. You can include a comma-separated list of as many tags as you need.

#### `MapTo` Tag

`MapTo[T]` adds code to transform the stream to a new stream of a different type by transforming its element using
a function you provide. `MapTo[T]` can be used more than once: 

````go
// +gen Plumbing:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

Each tag creates a corresponding `MapToFred`, `MapToJim` etc function. These functions return the corresponding
`FredSet`, `JimSet` etc values. A mapping function transforms an individual `Example` item to its `Fred` equivalent.

In addition, a corresponding `FlatMapToFred`, `FlatMapToJim` etc function is generated.
A mapping function transforms an individual `Example` item to zero or more `Fred` equivalents.

#### Contents:

 * [Intro](README.md)
 * [Lists](List.md)
 * [Options](Option.md)
 * [Sets](Set.md)
 * [Joint Lists With Options and/or Sets](Unified.md)
 * **Plumbing functions**
