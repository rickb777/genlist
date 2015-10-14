package list

const comparableFunctions = `
{{if .Type.Comparable}}
// These methods require {{.PName}} be comparable.

// Contains verifies that a given value is contained in {{.TName}}List.
func (list {{.TName}}List) Contains(value {{.PName}}) bool {
	for _, v := range list {
	    {{if .Type.Pointer}}
		if *v == *value {
			return true
		}
		{{else}}
		if v == value {
			return true
		}
		{{end}}
	}
	return false
}

// Count gives the number elements of {{.TName}}List that match a certain value.
func (list {{.TName}}List) Count(value {{.PName}}) (result int) {
	for _, v := range list {
	    {{if .Type.Pointer}}
		if *v == *value {
			result++
		}
		{{else}}
		if v == value {
			result++
		}
		{{end}}
	}
	return
}

// Distinct returns a new {{.TName}}List whose elements are unique.
func (list {{.TName}}List) Distinct() (result {{.TName}}List) {
	appended := make(map[{{.TName}}]bool)
	for _, v := range list {
	    {{if .Type.Pointer}}
		if !appended[*v] {
			result = append(result, v)
			appended[*v] = true
		}
		{{else}}
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
		{{end}}
	}
	return result
}
{{end}}
`
