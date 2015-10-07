package slice

import "github.com/clipperhouse/typewriter"

var count = &typewriter.Template{
	Name: "Count",
	Text: `
// Count gives the number elements of {{.SliceName}} that match a certain value.
func (rcv {{.SliceName}}) Count(value {{.Type}}) (result int) {
	for _, v := range rcv {
		if v == value {
			result++
		}
	}
	return
}
`}

var countBy = &typewriter.Template{
	Name: "CountBy",
	Text: `
// CountBy gives the number elements of {{.SliceName}} that return true for the passed func.
func (rcv {{.SliceName}}) CountBy(fn func({{.Type}}) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}
`}
