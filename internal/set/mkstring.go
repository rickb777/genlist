package set

const mkstringFunctions = `
// String implements the Stringer interface to render the set as a comma-separated array.
func (set {{.TName}}Set) String() string {
	return set.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (set {{.TName}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (set {{.TName}}Set) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(set)
	if l > 0 {
		sep := ""
		for v := range set {
			b.WriteString(sep)
			b.WriteString(fmt.Sprintf("%v", v))
			sep = mid
		}
	}
	b.WriteString(sfx)
	return b.String()
}

`
