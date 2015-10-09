package sequence

import "github.com/clipperhouse/typewriter"

var Comparable = typewriter.Template{
	Name: "Comparable",
	Text: `
// Distinct returns a new {{.Type}}List whose elements are unique.
func (v Some{{.Type}}) Distinct() (result {{.Type}}List) {
	result = append(result, v)
	return result
}

// Distinct returns a new {{.Type}}List whose elements are unique.
func (v no{{.Type}}) Distinct() {{.Type}}List {
	return
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
