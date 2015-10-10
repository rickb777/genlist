package genlist

import (
	"io"
	"github.com/rickb777/genlist/internal/option"
	"github.com/rickb777/typewriter"
)

func init() {
	err := typewriter.Register(NewOptionWriter())
	if err != nil {
		panic(err)
	}
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

	if err := writeBasicTemplate(w, tmpl, typ); err != nil {
		return err
	}

	_, err = sw.writeTemplateIfPossible(w, typ, option.Comparable)
	if err != nil {
		return err
	}

	for _, v := range tag.Values {
		err = sw.writeOne(w, typ, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sw *OptionWriter) writeTemplateIfPossible(w io.Writer, typ typewriter.Type, t typewriter.Template) (bool, error) {
	//	fmt.Printf("writeTemplateIfPossible %s\n", t.Name)
	if t.TypeConstraint.TryType(typ) == nil {
		v := typewriter.TagValue{}
		v.Name = t.Name
		tmpl, err := t.Parse()
		if err != nil {
			return false, err
		}
		err = writeTaggedTemplate(w, tmpl, typ, v)
		return err == nil, err
	}
	return false, nil
}

func (sw *OptionWriter) writeOne(w io.Writer, typ typewriter.Type, v typewriter.TagValue) error {
	tmpl, err := option.Templates.ByTagValue(typ, v)

	if err != nil {
		return err
	}

	return writeTaggedTemplate(w, tmpl, typ, v)
}
