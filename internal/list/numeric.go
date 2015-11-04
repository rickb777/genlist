package list

const numericFunctions = `
{{if .Type.Numeric}}
//-------------------------------------------------------------------------------------------------
// These methods are provided because {{.TName}} is numeric.

// Sum sums all elements in the list.
func (list {{.TName}}List) Sum() (result {{.PName}}) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (list {{.TName}}List) Mean() float64 {
	l := len(list)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length {{.TName}}List")
	}
	sum := {{if .Type.Ptr}}*(list.Sum()){{else}}list.Sum(){{end}}
	return float64(sum) / float64(l)
}

{{end}}
`
