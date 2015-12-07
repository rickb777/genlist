package list

const equalsFunctions = `
{{if .Type.FullyComparable}}
// Equals verifies that another {{.TName}}Collection has the same size and elements as this one. Also,
// because this collection is a sequence, the order must be the same.
// Omitted if {{.TName}} is not comparable.
func (list {{.TName}}List) Equals(other {{.TName}}Collection) bool {
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
			i++
		}
	})
	return eq
}

{{end}}

`
