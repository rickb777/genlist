package list

const ListMapToParamFunctions = `
{{if .TypeParameter.IsBasic}}
// MapTo{{.TypeParameter.LongName}} transforms {{.TName}}List to []{{.TypeParameter.Name}}.
func (list {{.TName}}List) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) []{{.TypeParameter}} {
	result := make([]{{.TypeParameter}}, 0, len(list))
	for _, v := range list {
		u := fn(v)
		result = append(result, {{.Addr}}u)
	}
	return result
}

// FlatMapTo{{.TypeParameter.LongName}} transforms {{.TName}}List to {{.TypeParameter.Name}}List, by repeatedly
// calling the supplied function and concatenating the results as a single flat list.
func (list {{.TName}}List) FlatMapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) []{{.TypeParameter}}) []{{.TypeParameter}} {
	result := make([]{{.TypeParameter}}, 0, len(list))
	for _, v := range list {
		u := fn(v)
		if len(u) > 0 {
			result = append(result, u...)
		}
	}
	return result
}

{{else}}
// MapTo{{.TypeParameter.LongName}} transforms {{.TName}}List to {{.TypeParameter.Name}}List.
func (list {{.TName}}List) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) {{.TypeParameter.Name}}Collection {
	result := make({{.TypeParameter.Name}}List, 0, len(list))
	for _, v := range list {
		u := fn(v)
		result = append(result, {{.Addr}}u)
	}
	return result
}

// FlatMapTo{{.TypeParameter.LongName}} transforms {{.TName}}List to {{.TypeParameter.Name}}List, by repeatedly
// calling the supplied function and concatenating the results as a single flat list.
func (list {{.TName}}List) FlatMapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter.Name}}Collection) {{.TypeParameter.Name}}Collection {
	result := make({{.TypeParameter.Name}}List, 0, len(list))
	for _, v := range list {
		u := fn(v)
		if u.NonEmpty() {
			result = append(result, (u.ToList())...)
		}
	}
	return result
}

{{end}}
`
