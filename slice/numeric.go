package slice

import "github.com/clipperhouse/typewriter"

var numeric = &typewriter.Template{
	Name: "Numeric",
	Text: `
// Sum sums {{.Type}} elements in {{.SliceName}}. See: http://clipperhouse.github.io/gen/#Sum
func (rcv {{.SliceName}}) Sum() (result {{.Type}}) {
	for _, v := range rcv {
		result += v
	}
	return
}

// Mean sums {{.SliceName}} over all elements and divides by len({{.SliceName}}). See: http://clipperhouse.github.io/gen/#Mean
func (rcv {{.SliceName}}) Mean() ({{.Type}}, error) {
	var result {{.Type}}

	l := len(rcv)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length {{.SliceName}}")
	}
	for _, v := range rcv {
		result += v
	}
	result = result / {{.Type}}(l)
	return result, nil
}
`,
	TypeConstraint: typewriter.Constraint{Numeric: true},
}
