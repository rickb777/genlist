package list

import "github.com/clipperhouse/typewriter"

var MapToT = &typewriter.Template{
	Name: "MapTo",
	Text: `
// MapTo{{.TypeParameter.LongName}} transforms a list of {{.TypeParameter}} from {{.ListName}}.
func (list {{.ListName}}) MapTo{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}List) {
	for _, v := range list {
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
