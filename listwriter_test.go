package genlist

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/rickb777/genlist/internal/list"
	"github.com/rickb777/typewriter"
)

var pkg *typewriter.Package

func init() {
	pkg = typewriter.NewPackage("dummy", "SomePackage")

	t1, err := pkg.Eval("int")

	if err != nil {
		panic(err)
	}

	t2, err := pkg.Eval("*rune")

	if err != nil {
		panic(err)
	}

	t1.Tags = typewriter.TagList{
		typewriter.Tag{
			Name: "list",
			Values: []typewriter.TagValue{
				{Name: "GroupBy", TypeParameters: []typewriter.Type{t2}},
				{Name: "Where", TypeParameters: nil},
			},
		},
	}

	pkg.Types = append(pkg.Types, t1)
}

func TestListWrite(t *testing.T) {
	for _, typ := range pkg.Types {
		var b bytes.Buffer

		sw := NewListWriter()

		b.WriteString(fmt.Sprintf("package %s\n\n", pkg.Name()))
		sw.Write(&b, typ)

		src := b.String()

		fset := token.NewFileSet()
		if _, err := parser.ParseFile(fset, "testwritelist.go", src, 0); err != nil {
			t.Error(err)
		}
	}
}
