package genlist

import (
	"io"
	"github.com/rickb777/genlist/internal/option"
	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewOptionWriter())
	if err != nil {
		panic(err)
	}
}

func OptionName(typ typewriter.Type) string {
	return "Optional" + typ.Name
}

type OptionWriter struct{}

func NewOptionWriter() *OptionWriter {
	return &OptionWriter{}
}

func (sw *OptionWriter) Name() string {
	return "Option"
}

func (sw *OptionWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (sw *OptionWriter) Write(w io.Writer, typ typewriter.Type) error {
	tag, found := typ.FindTag(sw)

	if !found {
		return nil
	}

	// start with the option template
	tmpl, err := option.Templates.ByTag(typ, tag)

	if err != nil {
		return err
	}

	m := option.Model{
		Type:      typ,
		OptionName: OptionName(typ),
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	for _, v := range tag.Values {
		var tp typewriter.Type

		if len(v.TypeParameters) > 0 {
			tp = v.TypeParameters[0]
		}

		m := option.Model{
			Type:          typ,
			OptionName:     OptionName(typ),
			TypeParameter: tp,
			TagValue:      v,
		}

		tmpl, err := option.Templates.ByTagValue(typ, v)

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	return nil
}
