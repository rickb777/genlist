package genlist

import (
	"io"
	"strings"
	"github.com/rickb777/genlist/internal/list"
	"github.com/rickb777/typewriter"
)

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
	return "List"
}

func (sw *ListWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
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

	// start with the list template
	tmpl, err := list.List.Parse()

	if err != nil {
		return err
	}

	if err := writeBasicTemplate(w, tmpl, typ); err != nil {
		return err
	}

	doneOrdered, err := sw.writeTemplateIfPossible(w, typ, list.Ordered)
	if !doneOrdered {
		_, err = sw.writeTemplateIfPossible(w, typ, list.NotOrdered)
	}
	if err != nil {
		return err
	}

	_, err = sw.writeTemplateIfPossible(w, typ, list.Numeric)
	if err != nil {
		return err
	}

	_, err = sw.writeTemplateIfPossible(w, typ, list.Comparable)
	if err != nil {
		return err
	}

	for _, v := range tag.Values {
		err = sw.writeOne(w, typ, v)
		if err != nil {
			return err
		}
	}

	if includeSortImplementation(tag.Values) {
		tmpl, err := list.SortImplementation.Parse()

		if err != nil {
			return err
		}

		if err := writeBasicTemplate(w, tmpl, typ); err != nil {
			return err
		}
	}

	return nil
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
		err = writeTaggedTemplate(w, tmpl, typ, v)
		return err == nil, err
	}
	return false, nil
}

func (sw *ListWriter) writeOne(w io.Writer, typ typewriter.Type, v typewriter.TagValue) error {
	tmpl, err := list.Templates.ByTagValue(typ, v)

	if err != nil {
		return err
	}

	return writeTaggedTemplate(w, tmpl, typ, v)
}

func includeSortImplementation(values []typewriter.TagValue) bool {
	for _, v := range values {
		if strings.HasPrefix(v.Name, "SortWith") {
			return true
		}
	}
	return false
}
