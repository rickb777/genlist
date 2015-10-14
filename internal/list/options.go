package list

const optionForList = `
// optionForList
{{if .Has.Option}}
// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (list {{.TName}}List) Find(fn func({{.PName}}) bool) Optional{{.TName}} {
	for _, v := range list {
		if fn(v) {
			//return Some{{.TName}}(v)
		}
	}
	return No{{.TName}}()
}

// HeadOption converts an option to a list of zero or one item
func (list {{.TName}}List) HeadOption() Optional{{.TName}} {
	if len(list) > 0 {
		return Some{{.TName}}(list[0])
	} else {
		return No{{.TName}}()
	}
}

{{end}}
`
