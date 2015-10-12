package option

var some = `
//-----------------------------------------------------------------------------

type Some{{.Type}} {{.Type}}

// cross-check
var _ Optional{{.Type}} = new(Some{{.Type}})

// Get returns the contained element if present. Otherwise it panics.
func (x Some{{.Type}}) Get() {{.Type}} {
	return {{.Type}}(x)
}

// GetOrElse returns the contained element if present. Otherwise it returns a default.
func (x Some{{.Type}}) GetOrElse(d func() {{.Type}}) {{.Type}} {
	return {{.Type}}(x)
}

// OrElse returns this option if present. Otherwise it returns the alternative.
func (x Some{{.Type}}) OrElse(alternative func() Optional{{.Type}}) Optional{{.Type}} {
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
func (x Some{{.Type}}) Find(predicate func({{.Type}}) bool) Optional{{.Type}} {
	if predicate({{.Type}}(x)) {
		return x
	}
	return No{{.Type}}()
}

// Exists verifies that one or more elements of Some{{.Type}} return true for the passed predicate.
func (x Some{{.Type}}) Exists(predicate func({{.Type}}) bool) bool {
	return predicate({{.Type}}(x))
}

// Forall verifies that all elements of Optional{{.Type}} return true for the passed predicate.
func (x Some{{.Type}}) Forall(predicate func({{.Type}}) bool) bool {
	return predicate({{.Type}}(x))
}

// Foreach iterates over Optional{{.Type}} and executes the passed function against each element.
func (x Some{{.Type}}) Foreach(fn func({{.Type}})) {
	fn({{.Type}}(x))
}

// Filter returns a Optional{{.Type}} whose elements return true for a predicate.
func (x Some{{.Type}}) Filter(predicate func({{.Type}}) bool) Optional{{.Type}} {
	if predicate({{.Type}}(x)) {
		return x
	}
	return No{{.Type}}()
}

{{if .Type.Comparable}}
// These methods require {{.Type}} be comparable.

// Contains verifies that a given value is contained in Optional{{.Type}}.
func (v Some{{.Type}}) Contains(value {{.Type}}) bool {
	if {{.Type}}(v) == value {
		return true
	}
	return false
}

// Count gives the number elements of Optional{{.Type}} that match a certain value.
func (v Some{{.Type}}) Count(value {{.Type}}) (result int) {
	if {{.Type}}(v) == value {
		result++
	}
	return
}
{{end}}

`
