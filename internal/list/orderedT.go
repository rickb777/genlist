package list

import "github.com/rickb777/typewriter"

var OrderedT = &typewriter.Template{
	Name: "Ordered",
	Text: `
// Min{{.TypeParameter.LongName}} selects the least value of {{.TypeParameter}} in {{.Type}}List.
// Returns error on {{.Type}}List with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (list {{.Type}}List) Min{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length {{.Type}}List")
		return
	}
	result = fn(list[0])
	if l > 1 {
		for _, v := range list[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// Max{{.TypeParameter.LongName}} selects the largest value of {{.TypeParameter}} in {{.Type}}List.
// Returns error on {{.Type}}List with no elements. See: http://clipperhouse.github.io/gen/#MaxCustom
func (list {{.Type}}List) Max{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length {{.Type}}List")
		return
	}
	result = fn(list[0])
	if l > 1 {
		for _, v := range list[1:] {
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
