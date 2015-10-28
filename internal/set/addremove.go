package set

const addRemoveFunctions = `
{{if .Has.Tag.Mutate}}
// Add mutates the set by adding elements to it.
// The receiver is modified and returned.
func (set {{.TName}}Set) Add(others ...{{.PName}}) {{.TName}}Set {
	for _, v := range others {
		set[{{.Ptr}}v] = struct{}{}
	}
	return set
}

// Remove mutates the set by removing elements from it.
// The receiver is modified and returned.
func (set {{.TName}}Set) Remove(unwanted ...{{.PName}}) {{.TName}}Set {
	for _, item := range unwanted {
		delete(set, {{.Ptr}}item)
	}
	return set
}

{{else}}
// Add creates a new set with elements added. This is similar to Union, but takes a slice of extra values.
// The receiver is not modified.
func (set {{.TName}}Set) Add(others ...{{.PName}}) {{.TName}}Set {
	added := New{{.TName}}Set()
	for item := range set {
		added[item] = struct{}{}
	}
	for _, item := range others {
		added[{{.Ptr}}item] = struct{}{}
	}
	return added
}

// Remove creates a new set with elements removed. This is similar to Difference, but takes a slice of unwanted values.
// The receiver is not modified.
func (set {{.TName}}Set) Remove(unwanted ...{{.PName}}) {{.TName}}Set {
	removed := New{{.TName}}Set()
	for item := range set {
		removed[item] = struct{}{}
	}
	for _, item := range unwanted {
		delete(removed, item)
	}
	return removed
}

{{end}}
`
