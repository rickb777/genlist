package list

import "github.com/clipperhouse/typewriter"

var List = &typewriter.Template{
	Name: "List",
	Text: `// {{.ListName}} is a slice of type {{.Type}}. Use it where you would use []{{.Type}}.
type {{.ListName}} []{{.Type}}

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list {{.ListName}}) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.ListName}}) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// IsEmpty tests whether {{.ListName}} is empty.
func (list {{.ListName}}) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether {{.ListName}} is empty.
func (list {{.ListName}}) NonEmpty() bool {
	return len(list) > 0
}

// Exists verifies that one or more elements of {{.ListName}} return true for the passed func.
func (list {{.ListName}}) Exists(fn func({{.Type}}) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.ListName}} return true for the passed func.
func (list {{.ListName}}) Forall(fn func({{.Type}}) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.ListName}} and executes the passed func against each element.
func (list {{.ListName}}) Foreach(fn func({{.Type}})) {
	for _, v := range list {
		fn(v)
	}
}

// Filter returns a new {{.ListName}} whose elements return true for func.
func (list {{.ListName}}) Filter(fn func({{.Type}}) bool) (result {{.ListName}}) {
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new {{.ListName}}s whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list {{.ListName}}) Partition(p func({{.Type}}) bool) (matching {{.ListName}}, others {{.ListName}}) {
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}

// Take returns a new {{.ListName}} containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.ListName}}) Take(n int) {{.ListName}} {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new {{.ListName}} without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.ListName}}) Drop(n int) {{.ListName}} {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new {{.ListName}} containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.ListName}}) TakeLast(n int) {{.ListName}} {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new {{.ListName}} without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list {{.ListName}}) DropLast(n int) {{.ListName}} {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0:l-n]
	}
}

// TakeWhile returns a new {{.ListName}} containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list {{.ListName}}) TakeWhile(p func({{.Type}}) bool) (result {{.ListName}}) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.ListName}} containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list {{.ListName}}) DropWhile(p func({{.Type}}) bool) (result {{.ListName}}) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Reverse returns a copy of {{.ListName}} with all elements in the reverse order.
func (list {{.ListName}}) Reverse() {{.ListName}} {
    numItems := len(list)
    result := make({{.ListName}}, numItems)
    last := numItems - 1
    for i, v := range list {
		result[last - i] = v
    }
    return result
}

// Shuffle returns a shuffled copy of {{.ListName}}, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (list {{.ListName}}) Shuffle() {{.ListName}} {
    numItems := len(list)
    result := make({{.ListName}}, numItems)
    copy(result, list)
    for i := 0; i < numItems; i++ {
        r := i + rand.Intn(numItems-i)
        result.Swap(i, r)
    }
    return result
}

// CountBy gives the number elements of {{.ListName}} that return true for the passed predicate.
func (list {{.ListName}}) CountBy(predicate func({{.Type}}) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.ListName}} containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list {{.ListName}}) MinBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// MaxBy returns an element of {{.ListName}} containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list {{.ListName}}) MaxBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// DistinctBy returns a new {{.ListName}} whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (list {{.ListName}}) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) (result {{.ListName}}) {
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
`,
}

// TODO aggregate, diff
// TODO PadTo, HeadOption, LastOption,
// TODO IndexOf, IndexWhere, LastIndexOf, LastIndexWhere
// TODO MkString, Equals, StartsWith, EndsWith, IteratorChan
// TODO FlatMap, Fold, FoldLeft, FoldRight
// TODO Find replaces First
