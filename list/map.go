package slice

import "github.com/clipperhouse/typewriter"

var mapToT = &typewriter.Template{
	Name: "MapTo",
	Text: `
// MapTo{{.TypeParameter.LongName}} transforms a slice of {{.TypeParameter}} from {{.ListName}}.
func (rcv {{.ListName}}) MapTo{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}List) {
	for _, v := range rcv {
		result = append(result, fn(v))
	}
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, but no constraints on that type
		{},
	},
}
