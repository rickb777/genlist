package slice

import "github.com/clipperhouse/typewriter"

var mean = &typewriter.Template{
	Name: "Mean",
	Text: `
// Mean sums {{.SliceName}} over all elements and divides by len({{.SliceName}}). See: http://clipperhouse.github.io/gen/#Mean
func (rcv {{.SliceName}}) Mean() ({{.Type}}, error) {
	var result {{.Type}}

	l := len(rcv)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length {{.SliceName}}")
	}
	for _, v := range rcv {
		result += v
	}
	result = result / {{.Type}}(l)
	return result, nil
}
`,
	TypeConstraint: typewriter.Constraint{Numeric: true},
}

var meanT = &typewriter.Template{
	Name: "Mean",
	Text: `
// Mean{{.TypeParameter.LongName}} sums {{.TypeParameter}} over all elements and divides by len({{.SliceName}}). See: http://clipperhouse.github.io/gen/#Mean
func (rcv {{.SliceName}}) Mean{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Mean[{{.TypeParameter}}] of zero-length {{.SliceName}}")
		return
	}
	for _, v := range rcv {
		result += fn(v)
	}
	result = result / {{.TypeParameter}}(l)
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be numeric
		{Numeric: true},
	},
}
