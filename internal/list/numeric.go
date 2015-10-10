package list

import "github.com/rickb777/typewriter"

var Numeric = typewriter.Template{
	Name: "Numeric",
	Text: `
// Sum sums {{.Type}} elements in {{.ListName}}. See: http://clipperhouse.github.io/gen/#Sum
func (list {{.ListName}}) Sum() (result {{.Type}}) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean sums {{.ListName}} over all elements and divides by len({{.ListName}}). See: http://clipperhouse.github.io/gen/#Mean
func (list {{.ListName}}) Mean() ({{.Type}}, error) {
	var result {{.Type}}

	l := len(list)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length {{.ListName}}")
	}
	for _, v := range list {
		result += v
	}
	result = result / {{.Type}}(l)
	return result, nil
}
`,
	TypeConstraint: typewriter.Constraint{Numeric: true},
}
