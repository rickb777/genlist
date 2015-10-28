package list

const mkstring = `
// String implements the Stringer interface to render the list as a comma-separated array.
func (list {{.TName}}List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (list {{.TName}}List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (list {{.TName}}List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", {{.Ptr}}v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", {{.Ptr}}v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}

`
