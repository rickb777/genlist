package list

const numericFunctions = `
{{if .Type.Numeric}}

// Sum sums {{.Type}} elements in {{.Type}}List. See: http://clipperhouse.github.io/gen/#Sum
func (list {{.Type}}List) Sum() (result {{.Type}}) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean sums {{.Type}}List over all elements and divides by len({{.Type}}List). See: http://clipperhouse.github.io/gen/#Mean
func (list {{.Type}}List) Mean() ({{.Type}}, error) {
	var result {{.Type}}

	l := len(list)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length {{.Type}}List")
	}
	for _, v := range list {
		result += v
	}
	result = result / {{.Type}}(l)
	return result, nil
}
{{end}}
`
