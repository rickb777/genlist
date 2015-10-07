package slice

import "github.com/clipperhouse/typewriter"

var sort = &typewriter.Template{
	Name: "Sort",
	Text: `
// Sort returns a new ordered {{.SliceName}}. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) Sort() {{.SliceName}} {
	result := make({{.SliceName}}, len(rcv))
	copy(result, rcv)
	sort.Sort(result)
	return result
}

// IsSorted reports whether {{.SliceName}} is sorted. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) IsSorted() bool {
	return sort.IsSorted(rcv)
}

// SortDesc returns a new reverse-ordered {{.SliceName}}. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) SortDesc() {{.SliceName}} {
	result := make({{.SliceName}}, len(rcv))
	copy(result, rcv)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether {{.SliceName}} is reverse-sorted. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(rcv))
}
`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}
