package list

const predicatedFunctions = `

// Filter returns a new {{.TName}}List whose elements return true for func.
func (list {{.TName}}List) Filter(fn func({{.PName}}) bool) {{.TName}}Seq {
	result := make({{.TName}}List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new {{.TName}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list {{.TName}}List) Partition(p func({{.PName}}) bool) (matching {{.TName}}List, others {{.TName}}List) {
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}

// CountBy gives the number elements of {{.TName}}List that return true for the passed predicate.
func (list {{.TName}}List) CountBy(predicate func({{.PName}}) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.TName}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list {{.TName}}List) MinBy(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the MinBy of an empty list.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// MaxBy returns an element of {{.TName}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list {{.TName}}List) MaxBy(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the MaxBy of an empty list.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if list[i] != list[m] && !less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// DistinctBy returns a new {{.TName}}List whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (list {{.TName}}List) DistinctBy(equal func({{.PName}}, {{.PName}}) bool) (result {{.TName}}List) {
Outer:
	for _, v := range list {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

`
