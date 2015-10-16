package golist

import (
	"github.com/rickb777/typewriter"
	"github.com/rickb777/golist/internal/option"
)

var coreOptionTemplate = &typewriter.Template{
	Name: "Option",
	Text: option.Optional,
}

var optionTemplates = typewriter.TemplateSlice{
	optionWithT,
	coreListTemplate,
}

var optionWithT = &typewriter.Template{
	Name: "With",
	Text: option.OptionMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

