package list

const WithParamFunctions = `

// Aggregate{{.TypeParameter.LongName}} iterates over {{.Type}}List, operating on each element while maintaining ‘state’.
func (list {{.Type}}List) Aggregate{{.TypeParameter.LongName}}(fn func({{.TypeParameter}}, {{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range list {
		result = fn(result, v)
	}
	return
}

{{if .TypeParameter.Comparable}}

// GroupBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}.
func (list {{.Type}}List) GroupBy{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.Type}}List {
	result := make(map[{{.TypeParameter}}]{{.Type}}List)
	for _, v := range list {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
{{end}}

{{if .TypeParameter.Numeric}}

// Sum{{.TypeParameter.LongName}} sums {{.Type}} over elements in {{.Type}}List. See: http://clipperhouse.github.io/gen/#Sum
func (list {{.Type}}List) Sum{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range list {
		result += fn(v)
	}
	return
}

// Mean{{.TypeParameter.LongName}} sums {{.TypeParameter}} over all elements and divides by len({{.Type}}List). See: http://clipperhouse.github.io/gen/#Mean
func (list {{.Type}}List) Mean{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Mean[{{.TypeParameter}}] of zero-length {{.Type}}List")
		return
	}
	for _, v := range list {
		result += fn(v)
	}
	result = result / {{.TypeParameter}}(l)
	return
}
{{end}}

{{if .TypeParameter.Ordered}}

// Min{{.TypeParameter.LongName}} selects the least value of {{.TypeParameter}} in {{.Type}}List.
// Returns error on {{.Type}}List with no elements.
func (list {{.Type}}List) Min{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length {{.Type}}List")
		return
	}
	result = fn(list[0])
	if l > 1 {
		for _, v := range list[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// Max{{.TypeParameter.LongName}} selects the largest value of {{.TypeParameter}} in {{.Type}}List.
// Returns error on {{.Type}}List with no elements.
func (list {{.Type}}List) Max{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length {{.Type}}List")
		return
	}
	result = fn(list[0])
	if l > 1 {
		for _, v := range list[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}
{{end}}
`
