package option

import "github.com/clipperhouse/typewriter"

var option = &typewriter.Template{
	Name: "Option",
	Text: `// {{.OptionName}} is an optional of type {{.Type}}. Use it where you want to be explicit about the presence	or absence of data.
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type {{.OptionName}} interface {
	Get() {{.Type}}
	GetOrElse(d func() {{.Type}}) {{.Type}}
	OrElse(alternative func() {{.OptionName}}) {{.OptionName}}
	Len() int
	IsEmpty() bool
	NonEmpty() bool
	Find(fn func({{.Type}}) bool) {{.OptionName}}
	Exists(fn func({{.Type}}) bool) bool
	Forall(fn func({{.Type}}) bool) bool
	Foreach(fn func({{.Type}}))
	Filter(fn func({{.Type}}) bool) (result {{.OptionName}})
}

//-----------------------------------------------------------------------------

type Some{{.Type}} {{.Type}}

// cross-check
var _ {{.OptionName}} = new(Some{{.Type}})

// Get returns the contained element if present. Otherwise it panics.
func (x Some{{.Type}}) Get() {{.Type}} {
	return {{.Type}}(x)
}

// GetOrElse returns the contained element if present. Otherwise it returns a default.
func (x Some{{.Type}}) GetOrElse(d func() {{.Type}}) {{.Type}} {
	return {{.Type}}(x)
}

// OrElse returns this option if present. Otherwise it returns the alternative.
func (x Some{{.Type}}) OrElse(alternative func() {{.OptionName}}) {{.OptionName}} {
	return x
}

// Len returns 1
func (x Some{{.Type}}) Len() int {
	return 1
}

// IsEmpty returns false.
func (x Some{{.Type}}) IsEmpty() bool {
	return false
}

// NonEmpty returns true.
func (x Some{{.Type}}) NonEmpty() bool {
	return true
}

// Find returns the contained value if it matches the predicate.
func (x Some{{.Type}}) Find(predicate func({{.Type}}) bool) {{.OptionName}} {
	if predicate({{.Type}}(x)) {
		return x
	}
	return No{{.Type}}()
}

// Exists verifies that one or more elements of Some{{.Type}} return true for the passed predicate.
func (x Some{{.Type}}) Exists(predicate func({{.Type}}) bool) bool {
	return predicate({{.Type}}(x))
}

// Forall verifies that all elements of {{.OptionName}} return true for the passed predicate.
func (x Some{{.Type}}) Forall(predicate func({{.Type}}) bool) bool {
	return predicate({{.Type}}(x))
}

// Foreach iterates over {{.OptionName}} and executes the passed function against each element.
func (x Some{{.Type}}) Foreach(fn func({{.Type}})) {
	fn({{.Type}}(x))
}

// Filter returns a {{.OptionName}} whose elements return true for a predicate.
func (x Some{{.Type}}) Filter(predicate func({{.Type}}) bool) {{.OptionName}} {
	if predicate({{.Type}}(x)) {
		return x
	}
	return No{{.Type}}()
}

//-----------------------------------------------------------------------------

type no{{.Type}} struct{}

// cross-check
var _ {{.OptionName}} = new(no{{.Type}})

func No{{.Type}}() {{.OptionName}} {
	return no{{.Type}}{}
}

// Get returns the contained element if present. Otherwise it panics.
func (x no{{.Type}}) Get() {{.Type}} {
	panic("Absent option - this indicates a promgramming error.")
}

// GetOrElse returns the contained element if present. Otherwise it returns a default.
func (x no{{.Type}}) GetOrElse(d func() {{.Type}}) {{.Type}} {
	return d()
}

// OrElse returns this option if present. Otherwise it returns the alternative.
func (x no{{.Type}}) OrElse(alternative func() {{.OptionName}}) {{.OptionName}} {
	return alternative()
}

// Len returns 0
func (x no{{.Type}}) Len() int {
	return 0
}

// IsEmpty returns true.
func (x no{{.Type}}) IsEmpty() bool {
	return true
}

// NonEmpty returns false.
func (x no{{.Type}}) NonEmpty() bool {
	return false
}

// Find can never succeed so returns No{{.Type}}.
func (x no{{.Type}}) Find(predicate func({{.Type}}) bool) {{.OptionName}} {
	return x
}

// Exists can never succeed so returns false.
func (x no{{.Type}}) Exists(predicate func({{.Type}}) bool) bool {
	return false
}

// Forall verifies that all elements of {{.OptionName}} return true for the passed predicate, which
// is always true.
func (x no{{.Type}}) Forall(predicate func({{.Type}}) bool) bool {
	return true
}

// Foreach iterates over {{.OptionName}} and executes the passed function against each element. There
// aren't any, so it does nothing.
func (x no{{.Type}}) Foreach(fn func({{.Type}})) {
	// does nothing
}

// Filter returns a {{.OptionName}} whose elements return true for a predicate.
func (x no{{.Type}}) Filter(predicate func({{.Type}}) bool) (result {{.OptionName}}) {
	return x
}

`,
}
