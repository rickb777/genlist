package slice

import "github.com/clipperhouse/typewriter"

var orderedT = &typewriter.Template{
	Name: "Ordered",
	Text: `
// Min{{.TypeParameter.LongName}} selects the least value of {{.TypeParameter}} in {{.ListName}}.
// Returns error on {{.ListName}} with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (rcv {{.ListName}}) Min{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length {{.ListName}}")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// Max{{.TypeParameter.LongName}} selects the largest value of {{.TypeParameter}} in {{.ListName}}.
// Returns error on {{.ListName}} with no elements. See: http://clipperhouse.github.io/gen/#MaxCustom
func (rcv {{.ListName}}) Max{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length {{.ListName}}")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be ordered
		{Ordered: true},
	},
}
