package golist

import (
	"github.com/rickb777/typewriter"
	"github.com/rickb777/golist/internal/set"
)

var coreSetTemplate = &typewriter.Template{
	Name: "Set",
	Text: set.Set,
}

var otherSetTemplates = typewriter.TemplateSlice{
	setMapToT,
	coreListTemplate,
}

var setMapToT = &typewriter.Template{
	Name: "MapTo",
	Text: set.SetMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

