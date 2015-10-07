package slice

import "github.com/clipperhouse/typewriter"

var distinct = &typewriter.Template{
	Name: "Distinct",
	Text: `
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


var distinctBy = &typewriter.Template{
	Name: "DistinctBy",
	Text: `
// DistinctBy returns a new {{.SliceName}} whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv {{.SliceName}}) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) (result {{.SliceName}}) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}
`}
