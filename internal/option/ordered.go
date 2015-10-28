package option

const ordered = `
{{if .Type.Ordered}}
// Min returns the minimum value of {{.TName}}List. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (o Optional{{.TName}}) Min() {{.PName}} {
	return o.Get()
}

// Max returns the maximum value of {{.TName}}List. In the case of multiple items being equally maximal,
// the first such element is returned. Panics if the collection is empty.
func (o Optional{{.TName}}) Max() {{.PName}} {
	return o.Get()
}

{{else}}
// Min returns an element of {{.TName}}List containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’. For ordered sequences, Min returns the first such element.
// Panics if the collection is empty.
func (o Optional{{.TName}}) Min(less func({{.PName}}, {{.PName}}) bool) {{.PName}} {
	return o.Get()
}

// Max returns an element of {{.TName}}List containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’. For ordered sequences, Max returns the first such element.
// Panics if the collection is empty.
func (o Optional{{.TName}}) Max(less func({{.PName}}, {{.PName}}) bool) {{.PName}} {
	return o.Get()
}

{{end}}
`
