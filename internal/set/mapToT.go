package set

const SetMapToParamFunctions = `
// Set:MapTo[{{.TypeParameter}}]

{{if .TypeParameter.IsBasic}}
// MapTo{{.TypeParameter.LongName}} transforms {{.TName}}Set to []{{.TypeParameter.Name}}.
func (set {{.TName}}Set) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) []{{.TypeParameter.Name}} {
	result := make([]{{.TypeParameter}}, 0, len(set))
	for v := range set {
		u := fn(v)
		result = append(result, {{.Addr}}u)
	}
	return result
}

// FlatMapTo{{.TypeParameter.LongName}} transforms {{.TName}}Set to []{{.TypeParameter}}, by repeatedly
// calling the supplied function and concatenating the results as a single flat slice.
func (set {{.TName}}Set) FlatMapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) []{{.TypeParameter}}) []{{.TypeParameter}} {
	result := make([]{{.TypeParameter}}, 0, len(set))
	for v := range set {
		u := fn(v)
		if len(u) > 0 {
			result = append(result, u...)
		}
	}
	return result
}

{{else}}
// MapTo{{.TypeParameter.LongName}} transforms {{.TName}}Set to {{.TypeParameter}}Set.
func (set {{.TName}}Set) MapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) {{.TypeParameter.Name}}Collection {
	result := make(map[{{.TypeParameter}}]struct{})
	for v := range set {
		u := fn(v)
		result[u] = struct{}{}
	}
	return {{.TypeParameter.Name}}Set(result)
}

// FlatMapTo{{.TypeParameter.LongName}} transforms {{.TName}}Set to {{.TypeParameter.Name}}Set, by
// calling the supplied function on each of the enclosed set elements, and returning a new set.
func (set {{.TName}}Set) FlatMapTo{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter.Name}}Collection) {{.TypeParameter.Name}}Collection {
	result := make(map[{{.TypeParameter}}]struct{})
	for a := range set {
		b := fn(a)
		b.Foreach(func (c {{.TypeParameter}}) {
			result[c] = struct{}{}
		})
	}
	return {{.TypeParameter.Name}}Set(result)
}

{{if .TypeParameter.Comparable}}
// GroupBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}.
// This method requires {{.TypeParameter.Name}} be comparable.
func (set {{.TName}}Set) GroupBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.TName}}Set {
	result := make(map[{{.TypeParameter}}]{{.TName}}Set)
	for v := range set {
		key := fn(v)
		group, exists := result[key]
		if !exists {
			group = New{{.TName}}Set()
		}
		group[v] = struct{}{}
		result[key] = group
	}
	return result
}

{{end}}
{{end}}
`
