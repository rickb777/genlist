package genlist

import (
	"github.com/rickb777/typewriter"
	"io"
)

type flags struct {
	Sequence bool
	Option   bool
	List     bool
}

// a convenience for passing values into templates; in MVC it'd be called a view model
type model struct {
	Type          typewriter.Type
	TName         string
	PName         string
	Ptr           string
	Deref         string
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
	Has           flags
}

func newModel(typ typewriter.Type, flags flags) model {
	p := ""
	d := "*"
	if typ.Pointer {
		p = "*"
		d = ""
	}
	return model{
		Type:  typ,
		TName: typ.Name,
		PName: typ.String(),
		Ptr:   p,
		Deref: d,
		Has:   flags,
	}
}

func writeBasicTemplate(w io.Writer, twTmpl *typewriter.Template, typ typewriter.Type, flags flags) error {
	flags.Sequence = true
	m := newModel(typ, flags)
	//	fmt.Printf("writeBasicTemplate\n  %+v\n", m)
	return writeTextTemplate(w, twTmpl, m)
}

func writeTaggedTemplate(w io.Writer, twTmpl *typewriter.Template, typ typewriter.Type, flags flags, tv typewriter.TagValue) error {
	var tp typewriter.Type

	if len(tv.TypeParameters) > 0 {
		tp = tv.TypeParameters[0]
	}

	m := newModel(typ, flags)
	m.TypeParameter = tp
	m.TagValue = tv

	//	fmt.Printf("writeTaggedTemplate\n  %+v\n", m)
	return writeTextTemplate(w, twTmpl, m)
}

func writeTextTemplate(w io.Writer, twTmpl *typewriter.Template, m model) error {
	tmpl, err := twTmpl.Parse()
	if err != nil {
		return err
	}
	return tmpl.Execute(w, m)
}

