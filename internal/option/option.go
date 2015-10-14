package option

import "github.com/rickb777/genlist/internal/sequence"

const Optional = sequence.Sequence + `
//-------------------------------------------------------------------------------------------------
// Optional{{.TName}} is an optional of type {{.PName}}. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type Optional{{.TName}} struct {
	x *{{.TName}}
}

var none{{.TName}} = Optional{{.TName}}{ nil }

func No{{.TName}}() Optional{{.TName}} {
	return none{{.TName}}
}

func Some{{.TName}}(x {{.PName}}) Optional{{.TName}} {
	{{if .Type.Pointer}}
	if x == nil {
		return none{{.TName}}
	}
	return Optional{{.TName}}{ x }
	{{else}}
	return Optional{{.TName}}{ &x }
	{{end}}
}

//-------------------------------------------------------------------------------------------------

func (o Optional{{.TName}}) Get() {{.PName}} {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return {{.Deref}}o.x
}

func (o Optional{{.TName}}) GetOrElse(d func() {{.PName}}) {{.PName}} {
	if o.IsEmpty() {
		return d()
	}
	return {{.Deref}}o.x
}

func (o Optional{{.TName}}) OrElse(alternative func() Optional{{.TName}}) Optional{{.TName}} {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//----- {{.TName}}Seq Methods -----

func (o Optional{{.TName}}) Len() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o Optional{{.TName}}) IsEmpty() bool {
	return o.x == nil
}

func (o Optional{{.TName}}) NonEmpty() bool {
	return o.x != nil
}

func (o Optional{{.TName}}) Find(predicate func({{.PName}}) bool) Optional{{.TName}} {
	if o.IsEmpty() {
		return o
	}
	if predicate({{.Deref}}o.x) {
		return o
	}
	return none{{.TName}}
}

func (o Optional{{.TName}}) Exists(predicate func({{.PName}}) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate({{.Deref}}o.x)
}

func (o Optional{{.TName}}) Forall(predicate func({{.PName}}) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate({{.Deref}}o.x)
}

func (o Optional{{.TName}}) Foreach(fn func({{.PName}})) {
	if o.NonEmpty() {
		fn({{.Deref}}o.x)
	}
}

func (o Optional{{.TName}}) Filter(predicate func({{.PName}}) bool) {{.TName}}Seq {
	return o.Find(predicate)
}

{{if .Type.Comparable}}
func (o Optional{{.TName}}) Contains(value {{.PName}}) bool {
	if *(o.x) == {{.Ptr}}value {
		return true
	}
	return false
}

func (o Optional{{.TName}}) Count(value {{.PName}}) (result int) {
	if o.Contains(value) {
		result++
	}
	return
}
{{end}}
{{if .Has.List}}
func (o Optional{{.TName}}) ToList() {{.TName}}List {
	if o.IsEmpty() {
		return {{.TName}}List{}
	}
	{{if .Type.Pointer}}
	return {{.TName}}List{ o.x }
	{{else}}
	return {{.TName}}List{ *(o.x) }
	{{end}}
}

{{end}}

`
