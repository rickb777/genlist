package list

import "github.com/clipperhouse/typewriter"

var Option = &typewriter.Template{
	Name: "Option",
	Text: `
// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (list {{.ListName}}) Find(fn func({{.Type}}) bool) Option{{.Type}} {
	for _, v := range list {
		if fn(v) {
			return Some{{.Type}}(v)
		}
	}
	err = errors.New("no {{.ListName}} elements return true for passed func")
	return No{{.Type}}
}
`}
