package sequence

const Sequence = `
{{if .Has.Sequence}}
//-------------------------------------------------------------------------------------------------
// {{.TName}}Seq is an interface for sequences of type {{.PName}}, including lists and options (where present).
type {{.TName}}Seq interface {
	{{.TName}}Collection

	// Len gets the size/length of the sequence - an alias for Size()
	Len() int

	//-------------------------------------------------------------------------
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() {{.PName}}

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() {{.PName}}

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() {{.TName}}Seq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() {{.TName}}Seq

{{if .Has.Option}}
	//-------------------------------------------------------------------------
	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func({{.PName}}) bool) Optional{{.TName}}
{{end}}

{{if .Has.List}}
	// Converts the sequence to a list. For lists, this is merely a type assertion.
	ToList() {{.TName}}List
{{end}}

{{if .Type.Comparable}}
	//-------------------------------------------------------------------------
	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if {{.TName}} is not comparable.
	Count(value {{.PName}}) int
{{if .Has.List}}

	// Distinct returns a new {{.TName}}Seq whose elements are all unique.
	// Omitted if {{.TName}} is not comparable.
	Distinct() {{.TName}}Seq
{{end}}
{{end}}
}


{{end}}
`
