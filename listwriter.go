package golist

import (
	"io"
	"github.com/rickb777/typewriter"
	"fmt"
)

const listName = "List"
const optionName = "Option"
const setName = "Set"

func init() {
	err := typewriter.Register(NewListWriter())
	if err != nil {
		panic(err)
	}
	err = typewriter.Register(NewOptionWriter())
	if err != nil {
		panic(err)
	}
	err = typewriter.Register(NewSetWriter())
	if err != nil {
		panic(err)
	}
}

type xWriter struct {
	name           string
	sequence       bool
	coreTemplate   *typewriter.Template
	otherTemplates typewriter.TemplateSlice
}

func NewListWriter() *xWriter {
	return &xWriter{listName, true, coreListTemplate, otherListTemplates}
}

func NewOptionWriter() *xWriter {
	return &xWriter{optionName, true, coreOptionTemplate, optionTemplates}
}

func NewSetWriter() *xWriter {
	return &xWriter{setName, false, coreSetTemplate, otherSetTemplates}
}

func (xw *xWriter) Name() string {
	return xw.name
}

func (xw *xWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (xw *xWriter) Write(w io.Writer, typ typewriter.Type) error {
	//	fmt.Printf("ListWriter.Write %s %+v\n", typ.String(), typ.Tags)
	tag, found := typ.FindTag(xw.Name())

	if !found {
		return nil
	}

	flags := setFlag(xw.name, &flags{})
	for _, v := range tag.Values {
		flags = setFlag(v.Name, flags)
	}

	// start with the list template
	flags.Sequence = xw.sequence
	if err := writeBasicTemplate(w, xw.coreTemplate, typ, *flags); err != nil {
		return err
	}

	flags.Sequence = false
	for _, v := range tag.Values {
		err := xw.writeForTag(w, typ, v, *flags)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(w, "// %s flags: %+v\n", xw.name, *flags)
	return nil
}

func (xw *xWriter) writeForTag(w io.Writer, typ typewriter.Type, v typewriter.TagValue, flags flags) error {
	tmpl, err := xw.otherTemplates.ByTagValue2(typ, v)

	if err != nil {
		return err
	}

	return writeTaggedTemplate(w, tmpl, typ, flags, v)
}

func setFlag(name string, flags *flags) *flags {
	switch name {
	case listName:   flags.List = true
	case optionName: flags.Option = true
	case setName:    flags.Set = true
	}
	return flags
}
