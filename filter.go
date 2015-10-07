package slice

import "github.com/clipperhouse/typewriter"

var filter = &typewriter.Template{
	Name: "Filter",
	Text: `
// Filter returns a new {{.SliceName}} whose elements return true for func.
func (rcv {{.SliceName}}) Filter(fn func({{.Type}}) bool) (result {{.SliceName}}) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
`}
