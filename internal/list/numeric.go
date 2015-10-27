package list

const numericFunctions = `
{{if .Type.Numeric}}
// These methods require {{.TName}} be numeric.

// Sum sums all elements in the list.
func (list {{.TName}}List) Sum() (result {{.PName}}) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (list {{.TName}}List) Mean() {{.PName}} {
	l := len(list)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length {{.TName}}List")
	}
	return list.Sum() / {{.PName}}(l)
}
{{end}}
`
