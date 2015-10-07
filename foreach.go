package slice

import "github.com/clipperhouse/typewriter"

var foreach = &typewriter.Template{
	Name: "Foreach",
	Text: `
// Foreach iterates over {{.SliceName}} and executes the passed func against each element.
func (rcv {{.SliceName}}) Foreach(fn func({{.Type}})) {
	for _, v := range rcv {
		fn(v)
	}
}
`}
