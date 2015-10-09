package slice

import (
	"io"
	"strings"
	"text/template"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewListWriter())
	if err != nil {
		panic(err)
	}
}

func ListName(typ typewriter.Type) string {
	return typ.Name + "List"
}

type ListWriter struct{}

func NewListWriter() *ListWriter {
	return &ListWriter{}
}

func (sw *ListWriter) Name() string {
	return "slice"
}

func (sw *ListWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (sw *ListWriter) writeTemplate(w io.Writer, typ typewriter.Type, v typewriter.TagValue, tmpl *template.Template) error {
	var tp typewriter.Type

	if len(v.TypeParameters) > 0 {
		tp = v.TypeParameters[0]
	}

	m := model{
		Type:          typ,
		ListName:     ListName(typ),
		TypeParameter: tp,
		TagValue:      v,
	}

	//	fmt.Printf("tmpl.Execute %s %s\n", typ.Name, v.Name)
	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	return nil
}

func (sw *ListWriter) writeOne(w io.Writer, typ typewriter.Type, v typewriter.TagValue) error {
	tmpl, err := templates.ByTagValue(typ, v)

	if err != nil {
		return err
	}

	return sw.writeTemplate(w, typ, v, tmpl)
}

func (sw *ListWriter) writeTemplateIfPossible(w io.Writer, typ typewriter.Type, t typewriter.Template) (bool, error) {
	//	fmt.Printf("writeTemplateIfPossible %s\n", t.Name)
	if t.TypeConstraint.TryType(typ) == nil {
		v := typewriter.TagValue{}
		v.Name = t.Name
		tmpl, err := t.Parse()
		if err != nil {
			return false, err
		}
		err = sw.writeTemplate(w, typ, v, tmpl)
		return err == nil, err
	}
	return false, nil
}

func (sw *ListWriter) Write(w io.Writer, typ typewriter.Type) error {
	tag, found := typ.FindTag(sw)

	if !found {
		return nil
	}

	//	fmt.Printf("\nWrite %s %+v\n", typ.Name, tag)
	if includeSortImplementation(tag.Values) {
		s := `// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

`
		w.Write([]byte(s))
	}

	// start with the slice template
	tmpl, err := slice.Parse()

	if err != nil {
		return err
	}

	m := model{
		Type:      typ,
		ListName: ListName(typ),
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	included := make(map[string]bool)

	doneOrdered, err := sw.writeTemplateIfPossible(w, typ, ordered)
	if doneOrdered {
		included[ordered.Name] = true
	} else {
		included[notOrdered.Name], err = sw.writeTemplateIfPossible(w, typ, notOrdered)
	}
	included[numeric.Name], err = sw.writeTemplateIfPossible(w, typ, numeric)
	included[comparable.Name], err = sw.writeTemplateIfPossible(w, typ, comparable)

	for _, v := range tag.Values {
		err = sw.writeOne(w, typ, v)
		if err != nil {
			return err
		}
		included[v.Name] = true
	}

	if includeSortImplementation(tag.Values) {
		tmpl, err := sortImplementation.Parse()

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	return nil
}

func includeSortImplementation(values []typewriter.TagValue) bool {
	for _, v := range values {
		if strings.HasPrefix(v.Name, "SortWith") {
			return true
		}
	}
	return false
}
