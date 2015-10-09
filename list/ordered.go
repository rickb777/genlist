package slice

import "github.com/clipperhouse/typewriter"

var ordered = typewriter.Template{
	Name: "Ordered",
	Text: `
// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (rcv {{.ListName}}) Less(i, j int) bool {
	return rcv[i] < rcv[j]
}

// Min returns the minimum value of {{.ListName}}. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
func (rcv {{.ListName}}) Min() (result {{.Type}}, err error) {
	if len(rcv) == 0 {
		err = errors.New("Cannot determine the Min of an empty slice.")
		return
	}
	result = rcv[0]
	for _, v := range rcv {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the maximum value of {{.ListName}}. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
func (rcv {{.ListName}}) Max() (result {{.Type}}, err error) {
	if len(rcv) == 0 {
		err = errors.New("Cannot determine the Max of an empty slice.")
		return
	}
	result = rcv[0]
	for _, v := range rcv {
		if v > result {
			result = v
		}
	}
	return
}

// Sort returns a new ordered {{.ListName}}.
func (rcv {{.ListName}}) Sort() {{.ListName}} {
	result := make({{.ListName}}, len(rcv))
	copy(result, rcv)
	sort.Sort(result)
	return result
}

// IsSorted reports whether {{.ListName}} is sorted.
func (rcv {{.ListName}}) IsSorted() bool {
	return sort.IsSorted(rcv)
}

// SortDesc returns a new reverse-ordered {{.ListName}}.
func (rcv {{.ListName}}) SortDesc() {{.ListName}} {
	result := make({{.ListName}}, len(rcv))
	copy(result, rcv)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether {{.ListName}} is reverse-sorted.
func (rcv {{.ListName}}) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(rcv))
}

`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}


// This is included when 'ordered' does not match its constraint.
var notOrdered = typewriter.Template{
	Name: "NotOrdered",
	Text: `
// Min returns an element of {{.ListName}} containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (rcv {{.ListName}}) Min(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("Cannot determine the Min of an empty slice.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// Max returns an element of {{.ListName}} containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (rcv {{.ListName}}) Max(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("Cannot determine the Max of an empty slice.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}


`,
}
