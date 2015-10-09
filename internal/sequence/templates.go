package sequence

import (
	"github.com/clipperhouse/typewriter"
)

// a convenience for passing values into templates; in MVC it'd be called a view model
type Model struct {
	Type          typewriter.Type
	SequenceName  string
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
}

var Templates = typewriter.TemplateSlice{
	Sequence,
}
