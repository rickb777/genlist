package sequence

const Sequence = `
// Sequence: Has {{.Has}}
{{if .Has.Sequence}}
// {{.TName}}Seq is an interface for sequences of type {{.PName}}, including lists and options (where present).
type {{.TName}}Seq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predictate supplied.
	Exists(predicate func({{.PName}}) bool) bool

	// Forall returns true if every element in the sequence matches the predictate supplied.
	Forall(predicate func({{.PName}}) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func({{.PName}}))

	// Filter returns a new {{.TName}}Seq whose elements return true for func.
	Filter(predicate func({{.PName}}) bool) (result {{.TName}}Seq)

{{if .Has.Option}}
	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func({{.PName}}) bool) Optional{{.TName}}
{{end}}

{{if .Has.List}}
	// Converts the sequence to a list. For lists, this is a no-op.
	ToList() {{.TName}}List
{{end}}

{{if .Type.Comparable}}
	// Contains tests whether a given value is present in the sequence.
	Contains(value {{.PName}}) bool

	// Count counts the number of times a given value occurs in the sequence.
	Count(value {{.PName}}) int
{{end}}
}

{{end}}
`
