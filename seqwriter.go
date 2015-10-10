package genlist

import (
	"io"
	"github.com/rickb777/genlist/internal/sequence"
	"github.com/rickb777/typewriter"
)

func init() {
	err := typewriter.Register(NewSequenceWriter())
	if err != nil {
		panic(err)
	}
}

type SequenceWriter struct{}

func NewSequenceWriter() *SequenceWriter {
	return &SequenceWriter{}
}

func (sw *SequenceWriter) Name() string {
	return "Seq"
}

func (sw *SequenceWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (sw *SequenceWriter) Write(w io.Writer, typ typewriter.Type) error {
	tag, found := typ.FindTag(sw)

	if !found {
		return nil
	}

	// start with the option template
	tmpl, err := sequence.Templates.ByTag(typ, tag)

	if err != nil {
		return err
	}

	m := sequence.Model{
		Type:      typ,
	}

	return tmpl.Execute(w, m)
}
