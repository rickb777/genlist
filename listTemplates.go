package genlist

import (
	"github.com/rickb777/typewriter"
"github.com/rickb777/genlist/internal/list"
)

var coreListTemplate = &typewriter.Template{
	Name: "List",
	Text: list.List,
}

var otherListTemplates = typewriter.TemplateSlice{
	withT,
	mapToT,
	sortWith,
	list.Option,
}


var sortWith = &typewriter.Template{
	Name: "SortWith",
	Text: list.SortWith,
}

var mapToT = &typewriter.Template{
	Name: "List",
	Text: list.ListMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

var withT = &typewriter.Template{
	Name: "With",
	Text: list.WithParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}
