package golist

import (
	"github.com/rickb777/typewriter"
	"github.com/rickb777/golist/internal/option"
)

var coreOptionTemplate = &typewriter.Template{
	Name: optionName,
	Text: option.Optional,
}

var optionTemplates = typewriter.TemplateSlice{
	optionMapToT,
	coreListTemplate,
	coreSetTemplate,
	corePlumbingTemplate,
}

var optionMapToT = &typewriter.Template{
	Name: "MapTo",
	Text: option.OptionMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

