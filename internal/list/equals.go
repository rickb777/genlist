package list

const equalsFunctions = `
{{if .Type.Comparable}}
// These methods require {{.PName}} be comparable.

// Equals verifies that one or more elements of {{.TName}}List return true for the passed func.
func (list {{.TName}}List) Equals(other {{.TName}}Seq) bool {
	if len(list) != other.Size() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a {{.PName}}) {
		if eq {
			v := list[i]
			if {{.Ptr}}v != {{.Ptr}}a {
				eq = false
			}
			i += 1
		}
	})
	return eq
}

{{end}}

`
