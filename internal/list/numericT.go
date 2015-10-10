package list

import "github.com/rickb777/typewriter"

var NumericT = &typewriter.Template{
	Name: "Numeric",
	Text: `
// Sum{{.TypeParameter.LongName}} sums {{.Type}} over elements in {{.ListName}}. See: http://clipperhouse.github.io/gen/#Sum
func (list {{.ListName}}) Sum{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range list {
		result += fn(v)
	}
	return
}

// Mean{{.TypeParameter.LongName}} sums {{.TypeParameter}} over all elements and divides by len({{.ListName}}). See: http://clipperhouse.github.io/gen/#Mean
func (list {{.ListName}}) Mean{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Mean[{{.TypeParameter}}] of zero-length {{.ListName}}")
		return
	}
	for _, v := range list {
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
