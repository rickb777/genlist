package list

import (
	"github.com/clipperhouse/typewriter"
)

// a convenience for passing values into templates; in MVC it'd be called a view model
type Model struct {
	Type      typewriter.Type
	ListName string
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
}

var Templates = typewriter.TemplateSlice{
	AggregateT,
	NumericT,
	OrderedT,
	GroupByT,
	MapToT,
	SortWith,

	Options,

	SortImplementation,
}
