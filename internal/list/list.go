package list

import "github.com/rickb777/typewriter"

var List = &typewriter.Template{
	Name: "List",
	Text: `// {{.Type}}List is a slice of type {{.Type}}. Use it where you would use []{{.Type}}.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// See e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq

type {{.Type}}List []{{.Type}}

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list {{.Type}}List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.Type}}List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// IsEmpty tests whether {{.Type}}List is empty.
func (list {{.Type}}List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether {{.Type}}List is empty.
func (list {{.Type}}List) NonEmpty() bool {
	return len(list) > 0
}

// Exists verifies that one or more elements of {{.Type}}List return true for the passed func.
func (list {{.Type}}List) Exists(fn func({{.Type}}) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.Type}}List return true for the passed func.
func (list {{.Type}}List) Forall(fn func({{.Type}}) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.Type}}List and executes the passed func against each element.
func (list {{.Type}}List) Foreach(fn func({{.Type}})) {
	for _, v := range list {
		fn(v)
	}
}

// Filter returns a new {{.Type}}List whose elements return true for func.
func (list {{.Type}}List) Filter(fn func({{.Type}}) bool) (result {{.Type}}List) {
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new {{.Type}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list {{.Type}}List) Partition(p func({{.Type}}) bool) (matching {{.Type}}List, others {{.Type}}List) {
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}

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

// Reverse returns a copy of {{.Type}}List with all elements in the reverse order.
func (list {{.Type}}List) Reverse() {{.Type}}List {
    numItems := len(list)
    result := make({{.Type}}List, numItems)
    last := numItems - 1
    for i, v := range list {
		result[last - i] = v
    }
    return result
}

// Shuffle returns a shuffled copy of {{.Type}}List, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (list {{.Type}}List) Shuffle() {{.Type}}List {
    numItems := len(list)
    result := make({{.Type}}List, numItems)
    copy(result, list)
    for i := 0; i < numItems; i++ {
        r := i + rand.Intn(numItems-i)
        result.Swap(i, r)
    }
    return result
}

// CountBy gives the number elements of {{.Type}}List that return true for the passed predicate.
func (list {{.Type}}List) CountBy(predicate func({{.Type}}) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.Type}}List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list {{.Type}}List) MinBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// MaxBy returns an element of {{.Type}}List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list {{.Type}}List) MaxBy(less func({{.Type}}, {{.Type}}) bool) (result {{.Type}}, err error) {
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

// DistinctBy returns a new {{.Type}}List whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (list {{.Type}}List) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) (result {{.Type}}List) {
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

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list {{.Type}}List) ToList() {{.Type}}List {
	return list
}
`,
}

// TODO aggregate, diff
// TODO PadTo, HeadOption, LastOption,
// TODO IndexOf, IndexWhere, LastIndexOf, LastIndexWhere
// TODO MkString, Equals, StartsWith, EndsWith, IteratorChan
// TODO FlatMap, Fold, FoldLeft, FoldRight
// TODO Find replaces First
