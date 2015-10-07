package slice

import "github.com/clipperhouse/typewriter"

var contains = &typewriter.Template{
	Name: "COntains",
	Text: `
// Exists verifies that one or more elements of {{.SliceName}} return true for the passed func.
func (slice {{.SliceName}}) Contains(value {{.Type}}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
`}
