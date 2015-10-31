package set

const WithParamFunctions = `
// FoldLeft{{.TypeParameter.LongName}} applies a binary operator to a start value and all elements of this set, going left to right.
// Note: the result is well-defined only if the operator function is associative and commutative.
func (set {{.TName}}Set) FoldLeft{{.TypeParameter.LongName}}(zero {{.TypeParameter}}, fn func({{.TypeParameter}}, {{.PName}}) {{.TypeParameter}}) {{.TypeParameter}} {
	sum := zero
	for v := range set {
		sum = fn(sum, v)
	}
	return sum
}

// FoldRight{{.TypeParameter.LongName}} applies a binary operator to a start value and all elements of this set, going right to left.
// This is an alias for FoldLeft{{.TypeParameter.LongName}}.
// Note: the result is well-defined only if the operator function is associative and commutative.
func (set {{.TName}}Set) FoldRight{{.TypeParameter.LongName}}(zero {{.TypeParameter}}, fn func({{.TypeParameter}}, {{.PName}}) {{.TypeParameter}}) {{.TypeParameter}} {
	return set.FoldLeft{{.TypeParameter.LongName}}(zero, fn)
}

{{if .TypeParameter.Numeric}}
// Sum{{.TypeParameter.LongName}} sums {{.PName}} over elements in {{.TName}}Set.
// This method requires {{.PName}} be numeric.
func (set {{.TName}}Set) Sum{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for v := range set {
		result += fn(v)
	}
	return
}

// Mean{{.TypeParameter.LongName}} sums {{.TypeParameter}} over all elements and divides by len({{.TName}}Set).
// This method requires {{.PName}} be numeric.
// Panics if there are no elements.
func (set {{.TName}}Set) Mean{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	l := len(set)
	if l == 0 {
		panic("Cannot determine the maximum of an empty set.")
		return
	}
	for v := range set {
		result += fn(v)
	}
	result = result / {{.TypeParameter}}(l)
	return
}
{{end}}

{{if .TypeParameter.Ordered}}
// MinBy{{.TypeParameter.LongName}} finds the first element which yields the smallest value measured by function fn.
// fn is usually called a projection or measuring function.
// Panics if there are no elements.
// This method requires {{.TypeParameter}} be ordered.
func (set {{.TName}}Set) MinBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.PName}}) {
	if len(set) == 0 {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m {{.TypeParameter}}
	first := true
	for v := range set {
		f := fn(v)
		if first {
			first = false
			result = {{.Ptr}}v
			m = f
		} else if m > f {
			result = {{.Ptr}}v
			m = f
		}
	}
	return
}

// MaxBy{{.TypeParameter.LongName}} finds the first element which yields the largest value measured by function fn.
// fn is usually called a projection or measuring function.
// Panics if there are no elements.
// This method requires {{.TypeParameter}} be ordered.
func (set {{.TName}}Set) MaxBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.PName}}) {
	if len(set) == 0 {
		panic("Cannot determine the maximum of an empty set.")
	}
	var m {{.TypeParameter}}
	first := true
	for v := range set {
		f := fn(v)
		if first {
			first = false
			result = {{.Ptr}}v
			m = f
		} else if m < f {
			result = {{.Ptr}}v
			m = f
		}
	}
	return
}

{{end}}
`
