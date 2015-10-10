package option

import "github.com/rickb777/typewriter"

var Option = &typewriter.Template{
	Name: "Option",
	Text: `// Optional{{.Type}} is an optional of type {{.Type}}. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type Optional{{.Type}} interface {
	Get() {{.Type}}
	GetOrElse(d func() {{.Type}}) {{.Type}}
	OrElse(alternative func() Optional{{.Type}}) Optional{{.Type}}
	Len() int
	IsEmpty() bool
	NonEmpty() bool
	Find(fn func({{.Type}}) bool) Optional{{.Type}}
	Exists(fn func({{.Type}}) bool) bool
	Forall(fn func({{.Type}}) bool) bool
	Foreach(fn func({{.Type}}))
	Filter(fn func({{.Type}}) bool) (result Optional{{.Type}})
}

` + some + none,
}
