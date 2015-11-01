package list

import "github.com/rickb777/golist/internal/collection"

const List = collection.Collection + `
//-------------------------------------------------------------------------------------------------
// {{.TName}}List is a slice of type {{.PName}}. Use it where you would use []{{.PName}}.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type {{.TName}}List []{{.PName}}

//-------------------------------------------------------------------------------------------------
// New{{.TName}}List constructs a new list containing the supplied values, if any.
func New{{.TName}}List(values ...{{.PName}}) {{.TName}}List {
	list := make({{.TName}}List, len(values))
	for i, v := range values {
		list[i] = v
	}
	return list
}

{{if .Type.Underlying.IsBasic}}
// New{{.TName}}ListFrom{{.Type.Underlying.LongName}}s constructs a new {{.TName}}List from a []{{.Type.Underlying}}.
func New{{.TName}}ListFrom{{.Type.Underlying.LongName}}s(values []{{.Type.Underlying}}) {{.TName}}List {
	list := make({{.TName}}List, len(values))
	for i, v := range values {
		list[i] = {{.TName}}(v)
	}
	return list
}

{{end}}
// Build{{.TName}}ListFromChan constructs a new {{.TName}}List from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.TName}}ListFromChan(source <-chan {{.PName}}) {{.TName}}List {
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
mkstring + optionForList


// TODO diff
// TODO PadTo,
// TODO StartsWith, EndsWith, IteratorChan
// TODO Fold
