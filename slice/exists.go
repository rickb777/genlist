package slice

import "github.com/clipperhouse/typewriter"

var exists = &typewriter.Template{
	Name: "Exists",
	Text: `
// Exists verifies that one or more elements of {{.SliceName}} return true for the passed func.
func (slice {{.SliceName}}) Exists(fn func({{.Type}}) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}
`}
