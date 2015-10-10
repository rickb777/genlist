package sequence

import "github.com/rickb777/typewriter"

var Sequence = &typewriter.Template{
	Name: "Seq",
	Text: `// {{.SequenceName}} is an interface for sequences of type {{.Type}}.
type {{.SequenceName}} interface {
	Len() int
	IsEmpty() bool
	NonEmpty() bool
	//Find(fn func({{.Type}}) bool) Optional{{.Type}}
	Exists(fn func({{.Type}}) bool) bool
	Forall(fn func({{.Type}}) bool) bool
	Foreach(fn func({{.Type}}))
	//Filter(fn func({{.Type}}) bool) (result {{.SequenceName}})
	ToList() {{.Type}}List
}

// ToList converts an option to a list of zero or one item
func (x Some{{.Type}}) ToList() {{.Type}}List {
	return {{.Type}}List{ {{.Type}}(x) }
}

// ToList converts an option to a list of zero or one item
func (x no{{.Type}}) ToList() {{.Type}}List {
	return {{.Type}}List{}
}

// HeadOption converts an option to a list of zero or one item
func (list {{.Type}}List) HeadOption() Optional{{.Type}} {
	if len(list) == 0 {
		return Some{{.Type}}(list[0])
	} else {
		return no{{.Type}}{}
	}
}

`,
}
