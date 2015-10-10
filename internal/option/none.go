package option

var none = `
//-----------------------------------------------------------------------------

type no{{.Type}} struct{}

// cross-check
var _ Optional{{.Type}} = new(no{{.Type}})

func No{{.Type}}() Optional{{.Type}} {
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
func (x no{{.Type}}) OrElse(alternative func() Optional{{.Type}}) Optional{{.Type}} {
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
func (x no{{.Type}}) Find(predicate func({{.Type}}) bool) Optional{{.Type}} {
	return x
}

// Exists can never succeed so returns false.
func (x no{{.Type}}) Exists(predicate func({{.Type}}) bool) bool {
	return false
}

// Forall verifies that all elements of Optional{{.Type}} return true for the passed predicate, which
// is always true.
func (x no{{.Type}}) Forall(predicate func({{.Type}}) bool) bool {
	return true
}

// Foreach iterates over Optional{{.Type}} and executes the passed function against each element. There
// aren't any, so it does nothing.
func (x no{{.Type}}) Foreach(fn func({{.Type}})) {
	// does nothing
}

// Filter returns a Optional{{.Type}} whose elements return true for a predicate.
func (x no{{.Type}}) Filter(predicate func({{.Type}}) bool) (result Optional{{.Type}}) {
	return x
}

{{if .Type.Comparable}}
// Contains verifies that a given value is contained in Optional{{.Type}}.
func (v no{{.Type}}) Contains(value {{.Type}}) bool {
	return false
}

// Count gives the number elements of Optional{{.Type}} that match a certain value.
func (v no{{.Type}}) Count(value {{.Type}}) int {
	return 0
}

{{end}}
`
