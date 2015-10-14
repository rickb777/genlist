package genlist

import (
	"io"
	"github.com/rickb777/typewriter"
)

const listName = "List"

func init() {
	err := typewriter.Register(NewListWriter())
	if err != nil {
		panic(err)
	}
}

type ListWriter struct{}

func NewListWriter() *ListWriter {
	return &ListWriter{}
}

func (sw *ListWriter) Name() string {
	return listName
}

func (sw *ListWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (sw *ListWriter) Write(w io.Writer, typ typewriter.Type) error {
	//	fmt.Printf("ListWriter.Write %s %+v\n", typ.String(), typ.Tags)
	tag, found := typ.FindTag(sw.Name())

	if !found {
		return nil
	}

	flags := flags{List: true}
	for _, v := range tag.Values {
		if v.Name == optionName {
			flags.Option = true
		}
	}

	// start with the list template
	if err := writeBasicTemplate(w, coreListTemplate, typ, flags); err != nil {
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

func (sw *ListWriter) writeForTag(w io.Writer, typ typewriter.Type, v typewriter.TagValue, flags flags) error {
	tmpl, err := otherListTemplates.ByTagValue2(typ, v)

	if err != nil {
		return err
	}

	return writeTaggedTemplate(w, tmpl, typ, flags, v)
}
