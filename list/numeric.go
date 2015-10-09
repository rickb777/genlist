package slice

import "github.com/clipperhouse/typewriter"

var numeric = typewriter.Template{
	Name: "Numeric",
	Text: `
// Sum sums {{.Type}} elements in {{.ListName}}. See: http://clipperhouse.github.io/gen/#Sum
func (rcv {{.ListName}}) Sum() (result {{.Type}}) {
	for _, v := range rcv {
		result += v
	}
	return
}

// Mean sums {{.ListName}} over all elements and divides by len({{.ListName}}). See: http://clipperhouse.github.io/gen/#Mean
func (rcv {{.ListName}}) Mean() ({{.Type}}, error) {
	var result {{.Type}}

	l := len(rcv)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length {{.ListName}}")
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
