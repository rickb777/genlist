package genlist

import (
	"io"
	"github.com/rickb777/typewriter"
)

const optionName = "Option"

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
	return optionName
}

func (sw *OptionWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (sw *OptionWriter) Write(w io.Writer, typ typewriter.Type) error {
	//	fmt.Printf("OptionWriter.Write %s %+v\n", typ.String(), typ.Tags)
	tag, found := typ.FindTag(sw.Name())

	if !found {
		return nil
	}

	flags := flags{Option: true}
	for _, v := range tag.Values {
		if v.Name == listName {
			flags.List = true
		}
	}

	// start with the option template
	if err := writeBasicTemplate(w, coreOptionTemplate, typ, flags); err != nil {
		return err
	}

	for _, v := range tag.Values {
		err := sw.writeForTag(w, typ, v, flags)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sw *OptionWriter) writeForTag(w io.Writer, typ typewriter.Type, v typewriter.TagValue, flags flags) error {
	tmpl, err := optionTemplates.ByTagValue2(typ, v)

	if err != nil {
		return err
	}

	return writeTaggedTemplate(w, tmpl, typ, flags, v)
}
