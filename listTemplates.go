package golist

import (
	"github.com/rickb777/typewriter"
	"github.com/rickb777/golist/internal/list"
)

var coreListTemplate = &typewriter.Template{
	Name: "List",
	Text: list.List,
}

var otherListTemplates = typewriter.TemplateSlice{
	listSortWith,
	listMapToT,
	listWithT,
//	list.OptionForList,
	coreOptionTemplate,
}

var listSortWith = &typewriter.Template{
	Name: "SortWith",
	Text: list.SortWith,
}

var listMapToT = &typewriter.Template{
	Name: "MapTo",
	Text: list.ListMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

var listWithT = &typewriter.Template{
	Name: "With",
	Text: list.WithParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}
