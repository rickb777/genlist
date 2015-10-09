package slice

import "github.com/clipperhouse/typewriter"

var slice = &typewriter.Template{
	Name: "slice",
	Text: `// {{.SliceName}} is a slice of type {{.Type}}. Use it where you would use []{{.Type}}.
type {{.SliceName}} []{{.Type}}

// Len returns the number of items in the slice.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (rcv {{.SliceName}}) Len() int {
	return len(rcv)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (rcv {{.SliceName}}) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// IsEmpty tests whether {{.SliceName}} is empty.
func (slice {{.SliceName}}) IsEmpty() bool {
	return len(slice) == 0
}

// NonEmpty tests whether {{.SliceName}} is empty.
func (slice {{.SliceName}}) NonEmpty() bool {
	return len(slice) > 0
}

// Exists verifies that one or more elements of {{.SliceName}} return true for the passed func.
func (slice {{.SliceName}}) Exists(fn func({{.Type}}) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.SliceName}} return true for the passed func.
func (slice {{.SliceName}}) Forall(fn func({{.Type}}) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.SliceName}} and executes the passed func against each element.
func (slice {{.SliceName}}) Foreach(fn func({{.Type}})) {
	for _, v := range slice {
		fn(v)
	}
}

// Filter returns a new {{.SliceName}} whose elements return true for func.
func (rcv {{.SliceName}}) Filter(fn func({{.Type}}) bool) (result {{.SliceName}}) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new {{.SliceName}}s whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original slice.
func (slice {{.SliceName}}) Partition(p func({{.Type}}) bool) (matching {{.SliceName}}, others {{.SliceName}}) {
	for _, v := range slice {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}

// TakeWhile returns a new {{.SliceName}} containing the leading elements of the source slice. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (slice {{.SliceName}}) TakeWhile(p func({{.Type}}) bool) (result {{.SliceName}}) {
	for _, v := range slice {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new {{.SliceName}} containing the trailing elements of the source slice. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (slice {{.SliceName}}) DropWhile(p func({{.Type}}) bool) (result {{.SliceName}}) {
	adding := false
	for _, v := range slice {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Reverse returns a copy of {{.SliceName}} with all elements in the reverse order.
func (rcv {{.SliceName}}) Reverse() {{.SliceName}} {
    numItems := len(rcv)
    result := make({{.SliceName}}, numItems)
    last := numItems - 1
    for i, v := range rcv {
		result[last - i] = v
    }
    return result
}

// Shuffle returns a shuffled copy of {{.SliceName}}, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv {{.SliceName}}) Shuffle() {{.SliceName}} {
    numItems := len(rcv)
    result := make({{.SliceName}}, numItems)
    copy(result, rcv)
    for i := 0; i < numItems; i++ {
        r := i + rand.Intn(numItems-i)
        result.Swap(i, r)
    }
    return result
}

// CountBy gives the number elements of {{.SliceName}} that return true for the passed predicate.
func (rcv {{.SliceName}}) CountBy(predicate func({{.Type}}) bool) (result int) {
	for _, v := range rcv {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.SliceName}} containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (rcv {{.SliceName}}) MinBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// MaxBy returns an element of {{.SliceName}} containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (rcv {{.SliceName}}) MaxBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// DistinctBy returns a new {{.SliceName}} whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv {{.SliceName}}) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) (result {{.SliceName}}) {
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
// TODO PadTo, TakeWhile, DropWhile, Drop, HeadOption, LastOption,
// TODO IndexOf, IndexWhere, LastIndexOf, LastIndexWhere
// TODO MkString, Equals, StartsWith, EndsWith, IteratorChan
// TODO FlatMap, Fold, FoldLeft, FoldRight
// TODO Find replaces First
