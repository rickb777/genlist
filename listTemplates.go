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
	list.WithT,
	list.ListT,
	sortWith,
	list.Option,
}


var sortWith = &typewriter.Template{
	Name: "SortWith",
	Text: list.SortWith,
}
