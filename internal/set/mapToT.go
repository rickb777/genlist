package set

const SetMapToParamFunctions = `

// MapTo{{.TypeParameter.LongName}} transforms {{.TName}}Set to {{.TypeParameter}}Set.
func (set {{.TName}}Set) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) {{.TypeParameter}}Seq {
	result := make(map[{{.TName}}]struct{})
	for v := range set {
		u := fn(v)
		result[u] = struct{}{}
	}
	return result
}

// FlatMapTo{{.TypeParameter.LongName}} transforms {{.TName}}Set to {{.TypeParameter.Name}}Set, by
// calling the supplied function on each of the enclosed set elements, and returning a new set.
func (set {{.TName}}Set) FlatMapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter.Name}}Seq) (result {{.TypeParameter.Name}}Seq) {
	result := make(map[{{.TName}}]struct{})
	for a := range set {
		b := fn(a)
		b.Foreach(func (c {{.TypeParameter.Name}}) {
			result[c] = struct{}{}
		})
	}
	return result
}

`
