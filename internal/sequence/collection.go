package sequence

const Collection = `
{{if .Has.Collection}}
//-------------------------------------------------------------------------------------------------
// {{.TName}}Collection is an interface for collections of type {{.TName}}, including sets, lists and options (where present).
type {{.TName}}Collection interface {
	// Size gets the size/length of the collection.
	Size() int

	// IsEmpty returns true if the collection is empty.
	IsEmpty() bool

	// NonEmpty returns true if the collection is non-empty.
	NonEmpty() bool

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func({{.PName}}) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func({{.PName}}) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func({{.PName}}))

	// Iter sends all elements along a channel of type {{.TName}}. For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	Iter() <-chan {{.PName}}

	//-------------------------------------------------------------------------
	// Filter returns a new {{.TName}}Collection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func({{.PName}}) bool) (result {{.TName}}Collection)

	// Partition returns two new {{.TName}}Collections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func({{.PName}}) bool) (matching {{.TName}}Collection, others {{.TName}}Collection)

{{if .Type.Comparable}}
	//-------------------------------------------------------------------------

	// Equals verifies that another {{.TName}}Collection has the same type, size and elements as this one.
	// Omitted if {{.TName}} is not comparable.
	Equals(other {{.TName}}Collection) bool

	// Contains tests whether a given value is present in the collection.
	// Omitted if {{.TName}} is not comparable.
	Contains(value {{.PName}}) bool

{{end}}
{{if .Type.Numeric}}
	//-------------------------------------------------------------------------
	// Sum sums {{.PName}} elements.
	// Omitted if {{.TName}} is not numeric.
	Sum() {{.PName}}

	// Mean computes the arithmetic mean of all elements. Panics if the collection is empty.
	// Omitted if {{.TName}} is not numeric.
	Mean() {{.PName}}

{{end}}
	//-------------------------------------------------------------------------
	// String gets a string representation of the collection. "[" and "]" surround
	// a comma-separated list of the elements.
	String() string

	// MkString gets a string representation of the collection. "[" and "]" surround a list
	// of the elements joined by the separator you provide.
	MkString(sep string) string

	// MkString3 gets a string representation of the collection. 'pfx' and 'sfx' surround a list
	// of the elements joined by the 'mid' separator you provide.
	MkString3(pfx, mid, sfx string) string
}


//-------------------------------------------------------------------------------------------------
{{if .Type.Ordered}}
// {{.TName}}OrderedCollection is an interface for collections of ordered types.
type {{.TName}}OrderedCollection interface {
	// Min returns the minimum value of {{.TName}}List. In the case of multiple items being equally minimal,
	// the first such element is returned. Panics if the collection is empty.
	Min() {{.PName}}

	// Max returns the maximum value of {{.TName}}List. In the case of multiple items being equally maximal,
	// the first such element is returned. Panics if the collection is empty.
	Max() {{.PName}}
}

{{else}}
// {{.TName}}UnorderedCollection is an interface for collections of unordered types.
type {{.TName}}UnorderedCollection interface {
	// Min returns an element of {{.TName}}List containing the minimum value, when compared to other elements
	// using a specified comparator function defining ‘less’. For ordered sequences, Min returns the first such element.
	// Panics if the collection is empty.
	Min(less func({{.PName}}, {{.PName}}) bool) {{.PName}}

	// Max returns an element of {{.TName}}List containing the maximum value, when compared to other elements
	// using a specified comparator function defining ‘less’. For ordered sequences, Max returns the first such element.
	// Panics if the collection is empty.
	Max(less func({{.PName}}, {{.PName}}) bool) {{.PName}}
}


{{end}}
{{end}}
`
