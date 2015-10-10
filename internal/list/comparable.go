package list

import "github.com/rickb777/typewriter"

var Comparable = typewriter.Template{
	Name: "Comparable",
	Text: `
// Contains verifies that a given value is contained in {{.ListName}}.
func (list {{.ListName}}) Contains(value {{.Type}}) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of {{.ListName}} that match a certain value.
func (list {{.ListName}}) Count(value {{.Type}}) (result int) {
	for _, v := range list {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new {{.ListName}} whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (list {{.ListName}}) Distinct() (result {{.ListName}}) {
	appended := make(map[{{.Type}}]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
