package sequence

const Collection = `
{{if .Has.Collection}}
// {{.TName}}Collection is an interface for collections of type {{.TName}}, including sets, lists and options (where present).
type {{.TName}}Collection interface {
	// Size gets the size/length of the sequence.
	Size() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool
}

{{end}}
`

// TODO queue functions: PopHead, PopLast, PushHead, PushLast
