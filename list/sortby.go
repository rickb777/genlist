package slice

import "github.com/clipperhouse/typewriter"

var sortWith = &typewriter.Template{
	Name: "SortWith",
	Text: `
// SortWith returns a new ordered {{.ListName}}, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.ListName}}) SortWith(less func({{.Type}}, {{.Type}}) bool) {{.ListName}} {
	result := make({{.ListName}}, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSort{{.ListName}}(result, less, 0, n, maxDepth)
	return result
}

// IsSortedWith reports whether an instance of {{.ListName}} is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.ListName}}) IsSortedWith(less func({{.Type}}, {{.Type}}) bool) bool {
	n := len(rcv)
	for i := n - 1; i > 0; i-- {
		if less(rcv[i], rcv[i-1]) {
			return false
		}
	}
	return true
}

// SortWithDesc returns a new, descending-ordered {{.ListName}}, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.ListName}}) SortWithDesc(less func({{.Type}}, {{.Type}}) bool) {{.ListName}} {
	greater := func(a, b {{.Type}}) bool {
		return less(b, a)
	}
	return rcv.SortWith(greater)
}

// IsSortedDesc reports whether an instance of {{.ListName}} is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.ListName}}) IsSortedWithDesc(less func({{.Type}}, {{.Type}}) bool) bool {
	greater := func(a, b {{.Type}}) bool {
		return less(b, a)
	}
	return rcv.IsSortedWith(greater)
}
`}
