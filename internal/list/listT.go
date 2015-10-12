package list

const ListMapToParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms {{.Type}}List to {{.TypeParameter}}List.
func (list {{.Type}}List) MapTo{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}List) {
	for _, v := range list {
		result = append(result, fn(v))
	}
	return
}

`
