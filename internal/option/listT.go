package option

import "github.com/rickb777/typewriter"

var withT = &typewriter.Template{
	Name: "With",
	Text: generalParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

const generalParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms Optional{{.Type}} to Optional{{.TypeParameter}}.
func (v Some{{.Type}}) MapTo{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) Optional{{.TypeParameter}} {
	u := fn({{.Type}}(v))
	return Some{{.TypeParameter}}(u)
}

// MapTo{{.TypeParameter.LongName}} transforms Optional{{.Type}} to Optional{{.TypeParameter}}.
func (v no{{.Type}}) MapTo{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) Optional{{.TypeParameter}} {
	return no{{.TypeParameter}}{}
}

`
