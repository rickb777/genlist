package list

const orderedFunctions = `
{{if .Type.Ordered}}
// These methods require {{.PName}} be ordered.

// Min returns the minimum value of {{.TName}}List. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
func (list {{.TName}}List) Min() (result {{.PName}}, err error) {
	if len(list) == 0 {
		err = errors.New("Cannot determine the Min of an empty list.")
		return
	}
	result = list[0]
	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the maximum value of {{.TName}}List. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
func (list {{.TName}}List) Max() (result {{.PName}}, err error) {
	if len(list) == 0 {
		err = errors.New("Cannot determine the Max of an empty list.")
		return
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

// Min returns an element of {{.TName}}List containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (list {{.TName}}List) Min(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the Min of an empty list.")
		return
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

// Max returns an element of {{.TName}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (list {{.TName}}List) Max(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the Max of an empty list.")
		return
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
