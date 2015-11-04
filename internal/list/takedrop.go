package list

const takeDropFunctions = `
//-------------------------------------------------------------------------------------------------

// Take returns a new {{.TName}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.TName}}List) Take(n int) {{.TName}}List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new {{.TName}}List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.TName}}List) Drop(n int) {{.TName}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new {{.TName}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.TName}}List) TakeLast(n int) {{.TName}}List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new {{.TName}}List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.TName}}List) DropLast(n int) {{.TName}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0:l-n]
	}
}

// TakeWhile returns a new {{.TName}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list {{.TName}}List) TakeWhile(p func({{.PName}}) bool) (result {{.TName}}List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.TName}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list {{.TName}}List) DropWhile(p func({{.PName}}) bool) (result {{.TName}}List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

`
