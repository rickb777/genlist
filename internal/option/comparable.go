package option

const comparable = `
{{if .Type.Comparable}}
//-------------------------------------------------------------------------------------------------
// These methods require {{.PName}} be comparable.

// Equals verifies that one or more elements of {{.TName}}List return true for the passed func.
func (o Optional{{.TName}}) Equals(other {{.TName}}Collection) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Size() > 1 {
		return false
	}
	a := o.x
	s := other.ToSlice()
	b := s[0]
	return *a == {{.Ptr}}b
}

func (o Optional{{.TName}}) Contains(value {{.PName}}) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == {{.Ptr}}value
}

func (o Optional{{.TName}}) Count(value {{.PName}}) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new {{.TName}}Collection whose elements are all unique. For options, this simply returns the
// receiver.
// Omitted if {{.TName}} is not comparable.
func (o Optional{{.TName}}) Distinct() {{.TName}}Collection {
	return o
}

{{end}}
`
