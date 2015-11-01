package golist

import (
	"github.com/rickb777/typewriter"
	"io"
)

type flags struct {
	Collection bool
	List       bool
	Option     bool
	Set        bool
	Tag        map[string]bool
}

func newFlags() flags {
	flags := flags{}
	flags.Tag = make(map[string]bool)
	return flags
}

func (f flags) setFlag(name string) flags {
	switch name {
	case listName:   f.List = true
	case optionName: f.Option = true
	case setName:    f.Set = true
	default:         f.Tag[name] = true
	}
	return f
}

// a convenience for passing values into templates; in MVC it'd be called a view model
type model struct {
	Type          typewriter.Type
	TName         string
	PName         string
	Ptr           string
	Deref         string
	Addr          string
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
	Has           flags
}

func newModel(typ typewriter.Type, flags flags) model {
	p := ""
	d := "*"
	a := ""
	if typ.Pointer {
		p = "*"
		d = ""
		a = "&"
	}
	return model{
		Type:  typ,
		TName: typ.Name,
		PName: typ.String(),
		Ptr:   p,
		Deref: d,
		Addr:  a,
		Has:   flags,
	}
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

