package slice

import "github.com/clipperhouse/typewriter"

var partition = &typewriter.Template{
	Name: "Partition",
	Text: `
// Partition returns two new {{.SliceName}}s whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original slice.
func (slice {{.SliceName}}) Partition(p func({{.Type}}) bool) (matching {{.SliceName}}, others {{.SliceName}}) {
	for _, v := range slice {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}
`}
