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

func SequenceName(typ typewriter.Type) string {
	return typ.Name + "Seq"
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
		SequenceName: SequenceName(typ),
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

//	for _, v := range tag.Values {
//		var tp typewriter.Type
//
//		if len(v.TypeParameters) > 0 {
//			tp = v.TypeParameters[0]
//		}
//
//		m := sequence.Model{
//			Type:          typ,
//			SequenceName:  SequenceName(typ),
//			TypeParameter: tp,
//			TagValue:      v,
//		}
//
//		tmpl, err := sequence.Templates.ByTagValue(typ, v)
//
//		if err != nil {
//			return err
//		}
//
//		if err := tmpl.Execute(w, m); err != nil {
//			return err
//		}
//	}

	return nil
}
