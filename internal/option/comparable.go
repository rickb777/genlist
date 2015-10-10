package option

import "github.com/rickb777/typewriter"

var Comparable = typewriter.Template{
	Name: "Comparable",
	Text: `
// Contains verifies that a given value is contained in Optional{{.Type}}.
func (v Some{{.Type}}) Contains(value {{.Type}}) bool {
	if {{.Type}}(v) == value {
		return true
	}
	return false
}

// Count gives the number elements of Optional{{.Type}} that match a certain value.
func (v Some{{.Type}}) Count(value {{.Type}}) (result int) {
	if {{.Type}}(v) == value {
		result++
	}
	return
}

// Contains verifies that a given value is contained in Optional{{.Type}}.
func (v no{{.Type}}) Contains(value {{.Type}}) bool {
	return false
}

// Count gives the number elements of Optional{{.Type}} that match a certain value.
func (v no{{.Type}}) Count(value {{.Type}}) int {
	return 0
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
