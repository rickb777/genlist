package slice

import "github.com/clipperhouse/typewriter"

var forall = &typewriter.Template{
	Name: "Forall",
	Text: `
// Forall verifies that all elements of {{.SliceName}} return true for the passed func.
func (rcv {{.SliceName}}) Forall(fn func({{.Type}}) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}
`}
