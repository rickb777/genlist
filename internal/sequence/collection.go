package sequence

const Collection = `
{{if .Has.Collection}}
//-------------------------------------------------------------------------------------------------
// {{.TName}}Collection is an interface for collections of type {{.TName}}, including sets, lists and options (where present).
type {{.TName}}Collection interface {
	// Size gets the size/length of the sequence.
	Size() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func({{.PName}}) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func({{.PName}}) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func({{.PName}}))

	// Iter sends all elements along a channel of type {{.TName}}.
	// The first time it is used, order of the elements is not well defined. But the order is stable, which means
	// it will give the same order each subsequent time it is used.
	Iter() <-chan {{.PName}}

	//-------------------------------------------------------------------------
	// Filter returns a new {{.TName}}Collection whose elements return true for a predicate function.
	Filter(predicate func({{.PName}}) bool) (result {{.TName}}Collection)

	// Partition returns two new {{.TName}}Collections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func({{.PName}}) bool) (matching {{.TName}}Collection, others {{.TName}}Collection)

{{if .Type.Comparable}}
	//-------------------------------------------------------------------------
	// These methods require {{.TName}} be comparable.

	// Equals verifies that one or more elements of {{.TName}}Collection return true for the passed func.
	Equals(other {{.TName}}Collection) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if {{.TName}} is not comparable.
	Contains(value {{.PName}}) bool

{{end}}
{{if .Type.Numeric}}
	//-------------------------------------------------------------------------
	// Sum sums {{.PName}} elements.
	// Omitted if {{.TName}} is not numeric.
	Sum() {{.PName}}

	// Mean computes the arithmetic mean of all elements.
	// Panics if the list is empty.
	Mean() {{.PName}}

{{end}}
}


{{end}}
`

// TODO queue functions: PopHead, PopLast, PushHead, PushLast
