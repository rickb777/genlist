package list

const WithParamFunctions = `
//-------------------------------------------------------------------------------------------------
// List:With[{{.TypeParameter}}]

// FoldLeft{{.TypeParameter.LongName}} applies a binary operator to a start value and all elements of this list, going left to right.
func (list {{.TName}}List) FoldLeft{{.TypeParameter.LongName}}(zero {{.TypeParameter}}, fn func({{.TypeParameter}}, {{.PName}}) {{.TypeParameter}}) {{.TypeParameter}} {
	sum := zero
	for _, v := range list {
		sum = fn(sum, v)
	}
	return sum
}

// FoldRight{{.TypeParameter.LongName}} applies a binary operator to a start value and all elements of this list, going right to left.
func (list {{.TName}}List) FoldRight{{.TypeParameter.LongName}}(zero {{.TypeParameter}}, fn func({{.TypeParameter}}, {{.PName}}) {{.TypeParameter}}) {{.TypeParameter}} {
	sum := zero
	for i := len(list) - 1; i >= 0; i-- {
		sum = fn(sum, list[i])
	}
	return sum
}

{{if .TypeParameter.Comparable}}
// This methods require {{.PName}} be comparable.

// GroupBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}.
func (list {{.TName}}List) GroupBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.TName}}List {
	result := make(map[{{.TypeParameter}}]{{.TName}}List)
	for _, v := range list {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
{{end}}

{{if .TypeParameter.Numeric}}
// These methods require {{.PName}} be numeric.

// Sum{{.TypeParameter.LongName}} sums {{.PName}} over elements in {{.TName}}List.
func (list {{.TName}}List) Sum{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range list {
		result += fn(v)
	}
	return
}

// Mean{{.TypeParameter.LongName}} sums {{.TypeParameter}} over all elements and divides by len({{.TName}}List).
// Panics if there are no elements.
func (list {{.TName}}List) Mean{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the mean of an empty list.")
	}
	for _, v := range list {
		result += fn(v)
	}
	result = result / {{.TypeParameter}}(l)
	return
}
{{end}}

{{if .TypeParameter.Ordered}}
// These methods require {{.TypeParameter}} be ordered.

// MinBy{{.TypeParameter.LongName}} finds the first element which yields the smallest value measured by function fn.
// fn is usually called a projection or measuring function.
// Panics if there are no elements.
func (list {{.TName}}List) MinBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.PName}}) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
		return
	}
	result = list[0]
	if l > 1 {
		min := fn(result)
		for i := 1; i < l; i++ {
			v := list[i]
			f := fn(v)
			if min > f {
				min = f
				result = v
			}
		}
	}
	return
}

// MaxBy{{.TypeParameter.LongName}} finds the first element which yields the largest value measured by function fn.
// fn is usually called a projection or measuring function.
// Panics if there are no elements.
func (list {{.TName}}List) MaxBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.PName}}) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	result = list[0]
	if l > 1 {
		max := fn(result)
		for i := 1; i < l; i++ {
			v := list[i]
			f := fn(v)
			if max < f {
				max = f
				result = v
			}
		}
	}
	return
}

{{end}}
`
