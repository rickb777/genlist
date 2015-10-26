package list

import "github.com/rickb777/golist/internal/sequence"

const List = sequence.Collection + sequence.Sequence + `
//-------------------------------------------------------------------------------------------------
// {{.TName}}List is a slice of type {{.PName}}. Use it where you would use []{{.PName}}.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.TName}}List []{{.PName}}

//-------------------------------------------------------------------------------------------------
// Build{{.TName}}ListFrom constructs a new {{.TName}}List from a channel that supplies values
// until it is closed.
func Build{{.TName}}ListFrom(source <-chan {{.PName}}) {{.TName}}List {
	result := make({{.TName}}List, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------
` + headTail + sortable +
iterationFunctions + takeDropFunctions + predicatedFunctions +
equalsFunctions + comparableFunctions + numericFunctions + orderedFunctions +
mkstringFunctions + optionForList


// TODO diff
// TODO PadTo,
// TODO StartsWith, EndsWith, IteratorChan
// TODO Fold
