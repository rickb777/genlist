package list

const ListMapToParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms {{.TName}}List to {{.TypeParameter.Name}}List.
func (list {{.TName}}List) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter.Name}}List) {
	for _, v := range list {
	    {{if .Type.Pointer}}
		u := fn(v)
		result = append(result, &u)
		{{else}}
		result = append(result, fn(v))
		{{end}}
	}
	return
}

`
