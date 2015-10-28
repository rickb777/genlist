package golist

import (
	"github.com/rickb777/typewriter"
	"github.com/rickb777/golist/internal/plumbing"
)

var corePlumbingTemplate = &typewriter.Template{
	Name: "Plumbing",
	Text: plumbing.Plumbing,
}

var plumbingTemplates = typewriter.TemplateSlice{
	plumbingMapToT,
}

var plumbingMapToT = &typewriter.Template{
	Name: "MapTo",
	Text: plumbing.PlumbingMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

