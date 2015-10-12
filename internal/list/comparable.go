package list

const comparableFunctions = `
{{if .Type.Comparable}}
// These methods require {{.Type}} be comarable.

// Contains verifies that a given value is contained in {{.Type}}List.
func (list {{.Type}}List) Contains(value {{.Type}}) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of {{.Type}}List that match a certain value.
func (list {{.Type}}List) Count(value {{.Type}}) (result int) {
	for _, v := range list {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new {{.Type}}List whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (list {{.Type}}List) Distinct() (result {{.Type}}List) {
	appended := make(map[{{.Type}}]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}
{{end}}
`
