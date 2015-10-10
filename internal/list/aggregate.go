package list

import "github.com/rickb777/typewriter"

var AggregateT = &typewriter.Template{
	Name: "Aggregate",
	Text: `
// Aggregate{{.TypeParameter.LongName}} iterates over {{.Type}}List, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (list {{.Type}}List) Aggregate{{.TypeParameter.LongName}}(fn func({{.TypeParameter}}, {{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range list {
		result = fn(result, v)
	}
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, but no constraints on that type
		{},
	},
}
