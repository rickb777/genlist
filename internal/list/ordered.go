package list

const orderedFunctions = `
{{if .Type.Ordered}}
//-------------------------------------------------------------------------------------------------
// These methods are provided because {{.TName}} is ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (list {{.TName}}List) Min() (result {{.PName}}) {
	if len(list) == 0 {
		panic("Cannot determine the Min of an empty list.")
	}
	result = list[0]
	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the element with the maximum value. In the case of multiple items being equally maximal,
// the first such element is returned. Panics if the collection is empty.
func (list {{.TName}}List) Max() (result {{.PName}}) {
	if len(list) == 0 {
		panic("Cannot determine the Max of an empty list.")
	}
	result = list[0]
	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return
}

{{else}}
//-------------------------------------------------------------------------------------------------
// These methods are included when {{.TName}} is not ordered.

// Min returns the first element containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Panics if the collection is empty.
func (list {{.TName}}List) Min(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// Max returns the first element containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Panics if the collection is empty.
func (list {{.TName}}List) Max(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if list[i] != list[m] && !less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

{{end}}
`
