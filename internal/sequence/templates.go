package sequence

import (
	"github.com/rickb777/typewriter"
)

// a convenience for passing values into templates; in MVC it'd be called a view model
type Model struct {
	Type          typewriter.Type
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
}

var Templates = typewriter.TemplateSlice{
	Sequence,
}
