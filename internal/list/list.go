package list

import "github.com/rickb777/typewriter"

var List = &typewriter.Template{
	Name: "List",
	Text: `// {{.Type}}List is a slice of type {{.Type}}. Use it where you would use []{{.Type}}.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
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

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list {{.Type}}List) ToList() {{.Type}}List {
	return list
}
` + takeDropFunctions + predicatedFunctions + comparableFunctions + numericFunctions + orderedFunctions,
}

// TODO aggregate, diff
// TODO PadTo, HeadOption, LastOption,
// TODO IndexOf, IndexWhere, LastIndexOf, LastIndexWhere
// TODO MkString, Equals, StartsWith, EndsWith, IteratorChan
// TODO FlatMap, Fold, FoldLeft, FoldRight
// TODO Find replaces First
