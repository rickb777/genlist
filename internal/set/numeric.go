package set

const numericFunctions = `
{{if .Type.Numeric}}
// These methods require {{.TName}} be numeric.

// Sum sums all elements in the set.
func (set {{.TName}}Set) Sum() (result {{.PName}}) {
	for v := range set {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the set is empty.
func (set {{.TName}}Set) Mean() float64 {
	l := len(set)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length {{.TName}}Set")
	}
	sum := {{if .Type.Ptr}}*(set.Sum()){{else}}set.Sum(){{end}}
	return float64(sum) / float64(l)
}
{{end}}
`
