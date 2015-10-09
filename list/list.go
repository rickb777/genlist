package slice

import "github.com/clipperhouse/typewriter"

var slice = &typewriter.Template{
	Name: "slice",
	Text: `// {{.ListName}} is a slice of type {{.Type}}. Use it where you would use []{{.Type}}.
type {{.ListName}} []{{.Type}}

// Len returns the number of items in the slice.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (rcv {{.ListName}}) Len() int {
	return len(rcv)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (rcv {{.ListName}}) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// IsEmpty tests whether {{.ListName}} is empty.
func (slice {{.ListName}}) IsEmpty() bool {
	return len(slice) == 0
}

// NonEmpty tests whether {{.ListName}} is empty.
func (slice {{.ListName}}) NonEmpty() bool {
	return len(slice) > 0
}

// Exists verifies that one or more elements of {{.ListName}} return true for the passed func.
func (slice {{.ListName}}) Exists(fn func({{.Type}}) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.ListName}} return true for the passed func.
func (slice {{.ListName}}) Forall(fn func({{.Type}}) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.ListName}} and executes the passed func against each element.
func (slice {{.ListName}}) Foreach(fn func({{.Type}})) {
	for _, v := range slice {
		fn(v)
	}
}

// Filter returns a new {{.ListName}} whose elements return true for func.
func (rcv {{.ListName}}) Filter(fn func({{.Type}}) bool) (result {{.ListName}}) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new {{.ListName}}s whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original slice.
func (slice {{.ListName}}) Partition(p func({{.Type}}) bool) (matching {{.ListName}}, others {{.ListName}}) {
	for _, v := range slice {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}

// Take returns a new {{.ListName}} containing the leading n elements of the source slice.
// If n is greater than the size of the slice, the whole slice is returned.
func (slice {{.ListName}}) Take(n int) {{.ListName}} {
	if n > len(slice) {
		return slice
	} else {
		return slice[0:n]
	}
}

// Drop returns a new {{.ListName}} without the leading n elements of the source slice.
// If n is greater than the size of the slice, the whole slice is returned.
func (slice {{.ListName}}) Drop(n int) {{.ListName}} {
	l := len(slice)
	if n > l {
		return slice[l:]
	} else {
		return slice[n:]
	}
}

// TakeLast returns a new {{.ListName}} containing the trailing n elements of the source slice.
// If n is greater than the size of the slice, the whole slice is returned.
func (slice {{.ListName}}) TakeLast(n int) {{.ListName}} {
	l := len(slice)
	if n > l {
		return slice
	} else {
		return slice[l-n:]
	}
}

// DropLast returns a new {{.ListName}} without the trailing n elements of the source slice.
// If n is greater than the size of the slice, the whole slice is returned.
func (slice {{.ListName}}) DropLast(n int) {{.ListName}} {
	l := len(slice)
	if n > l {
		return slice[l:]
	} else {
		return slice[0:l-n]
	}
}

// TakeWhile returns a new {{.ListName}} containing the leading elements of the source slice. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (slice {{.ListName}}) TakeWhile(p func({{.Type}}) bool) (result {{.ListName}}) {
	for _, v := range slice {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.ListName}} containing the trailing elements of the source slice. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (slice {{.ListName}}) DropWhile(p func({{.Type}}) bool) (result {{.ListName}}) {
	adding := false
	for _, v := range slice {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Reverse returns a copy of {{.ListName}} with all elements in the reverse order.
func (rcv {{.ListName}}) Reverse() {{.ListName}} {
    numItems := len(rcv)
    result := make({{.ListName}}, numItems)
    last := numItems - 1
    for i, v := range rcv {
		result[last - i] = v
    }
    return result
}

// Shuffle returns a shuffled copy of {{.ListName}}, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv {{.ListName}}) Shuffle() {{.ListName}} {
    numItems := len(rcv)
    result := make({{.ListName}}, numItems)
    copy(result, rcv)
    for i := 0; i < numItems; i++ {
        r := i + rand.Intn(numItems-i)
        result.Swap(i, r)
    }
    return result
}

// CountBy gives the number elements of {{.ListName}} that return true for the passed predicate.
func (rcv {{.ListName}}) CountBy(predicate func({{.Type}}) bool) (result int) {
	for _, v := range rcv {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.ListName}} containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (rcv {{.ListName}}) MinBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("Cannot determine the MinBy of an empty slice.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// MaxBy returns an element of {{.ListName}} containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (rcv {{.ListName}}) MaxBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("Cannot determine the MaxBy of an empty slice.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// DistinctBy returns a new {{.ListName}} whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv {{.ListName}}) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) (result {{.ListName}}) {
Outer:
	for _, v := range rcv {
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
