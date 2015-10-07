package slice

import "github.com/clipperhouse/typewriter"

var ordered = &typewriter.Template{
	Name: "Ordered",
	Text: `
// Min returns the minimum value of {{.SliceName}}. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
func (rcv {{.SliceName}}) Min() (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
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

// Max returns the maximum value of {{.SliceName}}. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
func (rcv {{.SliceName}}) Max() (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Max of an empty slice")
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

`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}
