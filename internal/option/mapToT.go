package option

const OptionMapToParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms Optional{{.TName}} to Optional{{.TypeParameter}}.
func (o Optional{{.TName}}) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) {{.TypeParameter}}Seq {
	if o.IsEmpty() {
		return No{{.TypeParameter}}()
	}
	{{if .Type.Pointer}}
	u := fn(o.x)
	{{else}}
	u := fn(*(o.x))
	{{end}}
	return Some{{.TypeParameter}}({{.Addr}}u)
}

// FlatMapTo{{.TypeParameter.LongName}} transforms Optional{{.TName}} to Optional{{.TypeParameter.Name}}, by
// calling the supplied function on the enclosed instance, if any, and returning an option.
// The result is only defined if *both* the receiver is defined and the function returns a non-empty sequence.
func (o Optional{{.TName}}) FlatMapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter.Name}}Seq) (result {{.TypeParameter.Name}}Seq) {
	if o.IsEmpty() {
		return No{{.TypeParameter}}()
	}
	u := fn({{.Deref}}(o.x))
	if u.IsEmpty() {
		return No{{.TypeParameter}}()
	}
	return Some{{.TypeParameter}}(u.Head())
}

`
