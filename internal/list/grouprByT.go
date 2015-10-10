package list

import "github.com/rickb777/typewriter"

var GroupByT = &typewriter.Template{
	Name: "GroupBy",
	Text: `
// GroupBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}. See: http://clipperhouse.github.io/gen/#GroupBy
func (list {{.Type}}List) GroupBy{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.Type}}List {
	result := make(map[{{.TypeParameter}}]{{.Type}}List)
	for _, v := range list {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be comparable
		{Comparable: true},
	},
}
