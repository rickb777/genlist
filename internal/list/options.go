package list

const optionForList = `
// optionForList
{{if .Has.Option}}
// First returns the first element that returns true for the passed func. Returns error if no elements return true.
func (list {{.TName}}List) Find(fn func({{.PName}}) bool) Optional{{.TName}} {
	for _, v := range list {
		if fn(v) {
			//return Some{{.TName}}(v)
		}
	}
	return No{{.TName}}()
}

// HeadOption gets the first item in the list, provided there is one.
func (list {{.TName}}List) HeadOption() Optional{{.TName}} {
	l := len(list)
	if l > 0 {
		return Some{{.TName}}(list[0])
	} else {
		return No{{.TName}}()
	}
}

// TailOption gets the last item in the list, provided there is one.
func (list {{.TName}}List) LastOption() Optional{{.TName}} {
	l := len(list)
	if l > 0 {
		return Some{{.TName}}(list[l-1])
	} else {
		return No{{.TName}}()
	}
}

{{end}}
`
