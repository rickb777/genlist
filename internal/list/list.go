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
` + headTail + sortable +
iterationFunctions + takeDropFunctions + predicatedFunctions +
equalsFunctions + comparableFunctions + numericFunctions + orderedFunctions +
mkstringFunctions + optionForList


// TODO diff
// TODO PadTo,
// TODO StartsWith, EndsWith, IteratorChan
// TODO Fold
