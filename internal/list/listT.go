package list

import "github.com/rickb777/typewriter"

var listT = &typewriter.Template{
	Name: "List",
	Text: listParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

const listParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms {{.Type}}List to {{.TypeParameter}}List.
func (list {{.Type}}List) MapTo{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}List) {
	for _, v := range list {
		result = append(result, fn(v))
	}
	return
}

`
