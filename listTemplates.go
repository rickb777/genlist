package genlist

import (
	"github.com/rickb777/typewriter"
"github.com/rickb777/genlist/internal/list"
)

var listTemplates = typewriter.TemplateSlice{
	list.WithT,
	list.ListT,
	sortWith,
	list.Option,
}


var sortWith = &typewriter.Template{
	Name: "SortWith",
	Text: list.SortWith,
}
