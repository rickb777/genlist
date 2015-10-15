package list

const numericFunctions = `
{{if .Type.Numeric}}
// These methods require {{.PName}} be numeric.

// Sum sums {{.PName}} elements in {{.TName}}List.
func (list {{.TName}}List) Sum() (result {{.PName}}) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean sums {{.TName}}List over all elements and divides by len({{.TName}}List).
func (list {{.TName}}List) Mean() ({{.PName}}, error) {
	var result {{.PName}}

	l := len(list)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length {{.TName}}List")
	}
	for _, v := range list {
		result += v
	}
	result = result / {{.PName}}(l)
	return result, nil
}
{{end}}
`
