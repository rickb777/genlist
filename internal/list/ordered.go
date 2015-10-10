package list

import "github.com/rickb777/typewriter"

var Ordered = typewriter.Template{
	Name: "Ordered",
	Text: `
// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list {{.ListName}}) Less(i, j int) bool {
	return list[i] < list[j]
}

// Min returns the minimum value of {{.ListName}}. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
func (list {{.ListName}}) Min() (result {{.Type}}, err error) {
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

// Max returns the maximum value of {{.ListName}}. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
func (list {{.ListName}}) Max() (result {{.Type}}, err error) {
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

// Sort returns a new ordered {{.ListName}}.
func (list {{.ListName}}) Sort() {{.ListName}} {
	result := make({{.ListName}}, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether {{.ListName}} is sorted.
func (list {{.ListName}}) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered {{.ListName}}.
func (list {{.ListName}}) SortDesc() {{.ListName}} {
	result := make({{.ListName}}, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether {{.ListName}} is reverse-sorted.
func (list {{.ListName}}) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}


// This is included when 'ordered' does not match its constraint.
var NotOrdered = typewriter.Template{
	Name: "NotOrdered",
	Text: `
// Min returns an element of {{.ListName}} containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (list {{.ListName}}) Min(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// Max returns an element of {{.ListName}} containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (list {{.ListName}}) Max(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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


`,
}
