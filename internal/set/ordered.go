package set

const orderedFunctions = `
{{if .Type.Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods require {{.TName}} be ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// any such element is returned. Panics if the collection is empty.
func (set {{.TName}}Set) Min() (result {{.PName}}) {
	if len(set) == 0 {
		panic("Cannot determine the minimum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = {{.Ptr}}v
		} else if v < result {
			result = {{.Ptr}}v
		}
	}
	return
}

// Max returns the element with the maximum value. In the case of multiple items being equally maximal,
// any such element is returned. Panics if the collection is empty.
func (set {{.TName}}Set) Max() (result {{.PName}}) {
	if len(set) == 0 {
		panic("Cannot determine the maximum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = {{.Ptr}}v
		} else if v > result {
			result = {{.Ptr}}v
		}
	}
	return
}

{{else}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.TName}} is not ordered.

// Min returns an element containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Panics if the collection is empty.
func (set {{.TName}}Set) Min(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}) {
	l := len(set)
	if l == 0 {
		panic("Cannot determine the minimum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = {{.Ptr}}v
		} else if less(v, result) {
			result = {{.Ptr}}v
		}
	}
	return
}

// Max returns an element containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Panics if the collection is empty.
func (set {{.TName}}Set) Max(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}) {
	l := len(set)
	if l == 0 {
		panic("Cannot determine the maximum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = {{.Ptr}}v
		} else if less(result, v) {
			result = {{.Ptr}}v
		}
	}
	return
}
{{end}}
`
