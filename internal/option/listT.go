package option

const OptionMapToParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms Optional{{.TName}} to Optional{{.TypeParameter}}.
func (o Optional{{.TName}}) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) Optional{{.TypeParameter}} {
	if o.IsEmpty() {
		return No{{.TypeParameter}}()
	}
	{{if .Type.Pointer}}
	u := fn(o.x)
	{{else}}
	u := fn(*(o.x))
	{{end}}
	return Some{{.TypeParameter}}(u)
}

`
