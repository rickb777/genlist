package slice

import "github.com/clipperhouse/typewriter"

var comparable = &typewriter.Template{
	Name: "Comparable",
	Text: `
// Contains verifies that a given value is contained in {{.SliceName}}.
func (slice {{.SliceName}}) Contains(value {{.Type}}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of {{.SliceName}} that match a certain value.
func (rcv {{.SliceName}}) Count(value {{.Type}}) (result int) {
	for _, v := range rcv {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new {{.SliceName}} whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv {{.SliceName}}) Distinct() (result {{.SliceName}}) {
	appended := make(map[{{.Type}}]bool)
	for _, v := range rcv {
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
