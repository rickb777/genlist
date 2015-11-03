package option

const mkstring = `
//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the option as an array of one element.
func (o Optional{{.TName}}) String() string {
	return o.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (o Optional{{.TName}}) MkString(sep string) string {
	return o.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (o Optional{{.TName}}) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

`
