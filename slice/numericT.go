package slice

import "github.com/clipperhouse/typewriter"

var numericT = &typewriter.Template{
	Name: "Numeric",
	Text: `
// Sum{{.TypeParameter.LongName}} sums {{.Type}} over elements in {{.SliceName}}. See: http://clipperhouse.github.io/gen/#Sum
func (rcv {{.SliceName}}) Sum{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}

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
