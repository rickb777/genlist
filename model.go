package genlist

import (
"github.com/rickb777/typewriter"
"text/template"
	"io"
)

// a convenience for passing values into templates; in MVC it'd be called a view model
type model struct {
	Type      typewriter.Type
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
}

func writeBasicTemplate(w io.Writer, tmpl *template.Template, typ typewriter.Type) error {

	m := model{
		Type:          typ,
	}

	return tmpl.Execute(w, m)
}

func writeTaggedTemplate(w io.Writer, tmpl *template.Template, typ typewriter.Type, tv typewriter.TagValue) error {
	var tp typewriter.Type

	if len(tv.TypeParameters) > 0 {
		tp = tv.TypeParameters[0]
	}

	m := model{
		Type:          typ,
		TypeParameter: tp,
		TagValue:      tv,
	}

	return tmpl.Execute(w, m)
}

