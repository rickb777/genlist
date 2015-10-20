package option

import "github.com/rickb777/golist/internal/sequence"

const Optional = sequence.Collection + sequence.Sequence + `
//-------------------------------------------------------------------------------------------------
// Optional{{.TName}} is an optional of type {{.PName}}. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type Optional{{.TName}} struct {
	x *{{.TName}}
}

// shared none value
var none{{.TName}} = Optional{{.TName}}{ nil }

func No{{.TName}}() Optional{{.TName}} {
	return none{{.TName}}
}

func Some{{.TName}}(x {{.PName}}) Optional{{.TName}} {
	{{if .Type.Pointer}}
	if x == nil {
		return No{{.TName}}()
	}
	return Optional{{.TName}}{ x }
	{{else}}
	return Optional{{.TName}}{ &x }
	{{end}}
}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o Optional{{.TName}}) Head() {{.PName}} {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return {{.Deref}}(o.x)
}

// panics if option is empty
func (o Optional{{.TName}}) Last() {{.PName}} {
	return o.Head()
}

// panics if option is empty
func (o Optional{{.TName}}) Tail() {{.TName}}Seq {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return none{{.TName}}
}

// panics if option is empty
func (o Optional{{.TName}}) Init() {{.TName}}Seq {
	return o.Tail()
}

//-------------------------------------------------------------------------------------------------

func (o Optional{{.TName}}) Get() {{.PName}} {
	return o.Head()
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

//-------------------------------------------------------------------------------------------------

func (o Optional{{.TName}}) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o Optional{{.TName}}) Len() int {
	return o.Size()
}

func (o Optional{{.TName}}) IsEmpty() bool {
	return o.x == nil
}

func (o Optional{{.TName}}) NonEmpty() bool {
	return o.x != nil
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o Optional{{.TName}}) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

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

func (o Optional{{.TName}}) Partition(predicate func({{.PName}}) bool) ({{.TName}}Seq, {{.TName}}Seq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate({{.Deref}}o.x) {
		return o, none{{.TName}}
	}
	return none{{.TName}}, o
}

{{if .Type.Comparable}}
//-------------------------------------------------------------------------------------------------
// These methods require {{.PName}} be comparable.

// Equals verifies that one or more elements of {{.TName}}List return true for the passed func.
func (o Optional{{.TName}}) Equals(other {{.TName}}Seq) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Size() > 1 {
		return false
	}
	a := o.Head()
	b := other.Head()
	return {{.Ptr}}a == {{.Ptr}}b
}

func (o Optional{{.TName}}) Contains(value {{.PName}}) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == {{.Ptr}}value
}

func (o Optional{{.TName}}) Count(value {{.PName}}) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new {{.TName}}Seq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if {{.TName}} is not comparable.
func (o Optional{{.TName}}) Distinct() {{.TName}}Seq {
	return o
}

{{end}}
{{if .Type.Numeric}}
//-------------------------------------------------------------------------------------------------
// Sum sums {{.PName}} elements.
// Omitted if {{.TName}} is not numeric.
func (o Optional{{.TName}}) Sum() {{.PName}} {
	{{if .Type.Pointer}}
	if o.IsEmpty() {
		r := 0
		return &r
	}
	return o.x
	{{else}}
	if o.IsEmpty() {
		return 0
	}
	return *(o.x)
	{{end}}
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (o Optional{{.TName}}) Mean() {{.PName}} {
	if o.IsEmpty() {
		panic("Cannot compute the arithmetic mean of zero-length Optional{{.TName}}")
	}
	return o.Sum()
}

{{end}}
{{if .Has.List}}
func (o Optional{{.TName}}) ToList() {{.TName}}List {
	if o.IsEmpty() {
		return {{.TName}}List{}
	}
	return {{.TName}}List{ {{.Deref}}o.x }
}

{{end}}

//-------------------------------------------------------------------------------------------------
// String implements the Stringer interface to render the option as an array of one element.
func (o Optional{{.TName}}) String() string {
	return o.MkString(",")
}

// MkString concatenates the values as a string.
func (o Optional{{.TName}}) MkString(sep string) string {
	return o.MkString3("[", sep, "]")
}

// MkString3 concatenates the values as a string.
func (o Optional{{.TName}}) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

`
