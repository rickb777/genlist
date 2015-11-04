package list

const conversions = `
//-------------------------------------------------------------------------------------------------

// ToSlice gets all the list's elements in a plain slice. This is simply a type conversion and is hardly needed
// for lists, because the underlying type can be used directly also.
// It is part of the {{.TName}}Collection interface.
func (list {{.TName}}List) ToSlice() []{{.PName}} {
	return []{{.PName}}(list)
}

{{if .Type.Underlying.IsBasic}}
// To{{.Type.Underlying.LongName}}s gets all the elements in a []{{.Type.Underlying}}.
func (list {{.TName}}List) To{{.Type.Underlying.LongName}}s() []{{.Type.Underlying}} {
	slice := make([]{{.Type.Underlying}}, len(list))
	for i, v := range list {
		slice[i] = {{.Type.Underlying}}(v)
	}
	return slice
}

{{end}}
// ToList simply returns the list in this case, but is part of the Collection interface.
func (list {{.TName}}List) ToList() {{.TName}}List {
	return list
}

{{if .Has.Set}}
// ToSet gets all the list's elements in a {{.TName}}Set.
func (list {{.TName}}List) ToSet() {{.TName}}Set {
	set := make(map[{{.TName}}]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	return {{.TName}}Set(set)
}

{{end}}
`
