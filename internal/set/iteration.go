package set

const iterationFunctions = `
// Exists verifies that one or more elements of {{.TName}}Set return true for the passed func.
func (set {{.TName}}Set) Exists(fn func({{.PName}}) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.TName}}Set return true for the passed func.
func (set {{.TName}}Set) Forall(fn func({{.PName}}) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.TName}}Set and executes the passed func against each element.
// The order of the elements is not well defined but is probably repeatably stable until the set is changed.
func (set {{.TName}}Set) Foreach(fn func({{.PName}})) {
	for v := range set {
		fn(v)
	}
}

// Send sends all elements along a channel of type {{.TName}}.
// The order of the elements is not well defined but is probably repeatably stable until the set is changed.
func (set {{.TName}}Set) Send() <-chan {{.PName}} {
	ch := make(chan {{.PName}})
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

`
