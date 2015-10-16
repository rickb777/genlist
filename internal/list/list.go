package list

import "github.com/rickb777/golist/internal/sequence"

const List = sequence.Sequence + `
//-------------------------------------------------------------------------------------------------
// {{.TName}}List is a slice of type {{.PName}}. Use it where you would use []{{.PName}}.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.TName}}List []{{.PName}}

//-------------------------------------------------------------------------------------------------
` + sortable + `

// panics if list is empty
func (list {{.TName}}List) Head() {{.PName}} {
	return list[0]
}

// IsEmpty tests whether {{.TName}}List is empty.
func (list {{.TName}}List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether {{.TName}}List is empty.
func (list {{.TName}}List) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list {{.TName}}List) ToList() {{.TName}}List {
	return list
}
` +
iterationFunctions + takeDropFunctions + predicatedFunctions +
equalsFunctions + comparableFunctions + numericFunctions + orderedFunctions +
optionForList


// TODO diff
// TODO PadTo,
// TODO IndexOf, IndexWhere, LastIndexOf, LastIndexWhere
// TODO MkString, Equals, StartsWith, EndsWith, IteratorChan
// TODO FlatMap, Fold
