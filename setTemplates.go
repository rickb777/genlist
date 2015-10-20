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
	coreSetTemplate,
}
