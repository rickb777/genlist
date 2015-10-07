package slice

import (
	"github.com/clipperhouse/typewriter"
)

// a convenience for passing values into templates; in MVC it'd be called a view model
type model struct {
	Type      typewriter.Type
	SliceName string
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
}

var templates = typewriter.TemplateSlice{
	aggregateT,
	comparable,
	numeric,
	numericT,
	ordered,
	orderedT,
	first,
	comparableT,
	mapToT,

	sort,
	sortBy,

	sortImplementation,
	sortInterface,
}
