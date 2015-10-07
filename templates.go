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
	slice,
	size,

	aggregateT,
	forall,
	exists,
	average,
	averageT,
	count,
	distinct,
	distinctBy,
	foreach,
	first,
	groupByT,
	max,
	maxT,
	maxBy,
	min,
	minT,
	minBy,
	mapToT,
	single,
	sum,
	sumT,
	filter,

	sort,
	isSorted,
	sortDesc,
	isSortedDesc,

	sortBy,
	isSortedBy,
	sortByDesc,
	isSortedByDesc,

	sortImplementation,
	sortInterface,

    shuffle,
}
