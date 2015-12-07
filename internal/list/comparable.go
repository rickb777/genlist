package list

const comparableFunctions = `
{{if .Type.FullyComparable}}
//-------------------------------------------------------------------------------------------------
// These methods are provided because {{.PName}} is comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list {{.TName}}List) IndexOf(value {{.PName}}) int {
	for i, v := range list {
		if {{.Ptr}}v == {{.Ptr}}value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list {{.TName}}List) IndexOf2(value {{.PName}}, from int) int {
	for i, v := range list {
		if i >= from && {{.Ptr}}v == {{.Ptr}}value {
			return i
		}
	}
	return -1
}

// Contains verifies that a given value is contained in {{.TName}}List.
func (list {{.TName}}List) Contains(value {{.PName}}) bool {
	for _, v := range list {
		if {{.Ptr}}v == {{.Ptr}}value {
			return true
		}
	}
	return false
}

// Count gives the number elements of {{.TName}}List that match a certain value.
func (list {{.TName}}List) Count(value {{.PName}}) (result int) {
	for _, v := range list {
		if {{.Ptr}}v == {{.Ptr}}value {
			result++
		}
	}
	return
}

// Distinct returns a new {{.TName}}List whose elements are unique, retaining the original order.
func (list {{.TName}}List) Distinct() {{.TName}}Collection {
	result := make({{.TName}}List, 0)
	appended := make(map[{{.TName}}]bool)
	for _, v := range list {
		if !appended[{{.Ptr}}v] {
			result = append(result, v)
			appended[{{.Ptr}}v] = true
		}
	}
	return result
}
{{end}}
`
