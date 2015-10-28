package golist

import (
	"io"
	"github.com/rickb777/typewriter"
	"fmt"
)

const listName = "List"
const optionName = "Option"
const setName = "Set"
const plumbingName = "Plumbing"

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
	err = typewriter.Register(NewPlumbingWriter())
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

func NewPlumbingWriter() *xWriter {
	return &xWriter{plumbingName, false, corePlumbingTemplate, typewriter.TemplateSlice{}}
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

	flags := newFlags().setFlag(xw.name)
	for _, v := range tag.Values {
		flags = flags.setFlag(v.Name)
	}

	// enable collection and (possibly) sequence for the core template
	flags.Collection = true
	flags.Sequence = xw.sequence

	// start with the core template
	model := newModel(typ, flags)
	//	fmt.Printf("writeBasicTemplate\n  %+v\n", m)
	if err := writeTextTemplate(w, xw.coreTemplate, model); err != nil {
		return err
	}

	// only want to include collection and sequence once, so turn them off before processing tags
	flags.Collection = false
	flags.Sequence = false

	for _, v := range tag.Values {
		err := xw.writeForTag(w, typ, v, flags)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(w, "// %s flags: %+v\n", xw.name, flags)
	return nil
}

func (xw *xWriter) writeForTag(w io.Writer, typ typewriter.Type, v typewriter.TagValue, flags flags) error {
	tmpl, err := xw.otherTemplates.ByTagValue2(typ, v)

	if err != nil {
		return err
	}

	return writeTaggedTemplate(w, tmpl, typ, flags, v)
}
