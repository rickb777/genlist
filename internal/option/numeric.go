package option

const numeric = `
{{if .Type.Numeric}}
//-------------------------------------------------------------------------------------------------
// Sum sums {{.PName}} elements.
// Omitted if {{.TName}} is not numeric.
func (o Optional{{.TName}}) Sum() {{.PName}} {
	{{if .Type.Pointer}}
	if o.IsEmpty() {
		r := 0
		return &r
	}
	return o.x
	{{else}}
	if o.IsEmpty() {
		return 0
	}
	return *(o.x)
	{{end}}
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (o Optional{{.TName}}) Mean() float64 {
	if o.IsEmpty() {
		panic("Cannot compute the arithmetic mean of zero-length Optional{{.TName}}")
	}
	return float64(*(o.x))
}

{{end}}
`
