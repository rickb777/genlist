package list

const takeDropFunctions = `
// Take returns a new {{.Type}}List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) Take(n int) {{.Type}}List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new {{.Type}}List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) Drop(n int) {{.Type}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new {{.Type}}List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) TakeLast(n int) {{.Type}}List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new {{.Type}}List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.Type}}List) DropLast(n int) {{.Type}}List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0:l-n]
	}
}

// TakeWhile returns a new {{.Type}}List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list {{.Type}}List) TakeWhile(p func({{.Type}}) bool) (result {{.Type}}List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.Type}}List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list {{.Type}}List) DropWhile(p func({{.Type}}) bool) (result {{.Type}}List) {
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
