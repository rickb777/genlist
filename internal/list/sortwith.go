package list

import "github.com/rickb777/typewriter"

var SortWith = &typewriter.Template{
	Name: "SortWith",
	Text: `

//-----------------------------------------------------------------------------
// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.
//-----------------------------------------------------------------------------

// SortWith returns a new ordered {{.Type}}List, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (list {{.Type}}List) SortWith(less func({{.Type}}, {{.Type}}) bool) {{.Type}}List {
	result := make({{.Type}}List, len(list))
	copy(result, list)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSort{{.Type}}List(result, less, 0, n, maxDepth)
	return result
}

// IsSortedWith reports whether an instance of {{.Type}}List is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (list {{.Type}}List) IsSortedWith(less func({{.Type}}, {{.Type}}) bool) bool {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		if less(list[i], list[i-1]) {
			return false
		}
	}
	return true
}

// SortWithDesc returns a new, descending-ordered {{.Type}}List, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (list {{.Type}}List) SortWithDesc(less func({{.Type}}, {{.Type}}) bool) {{.Type}}List {
	greater := func(a, b {{.Type}}) bool {
		return less(b, a)
	}
	return list.SortWith(greater)
}

// IsSortedDesc reports whether an instance of {{.Type}}List is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (list {{.Type}}List) IsSortedWithDesc(less func({{.Type}}, {{.Type}}) bool) bool {
	greater := func(a, b {{.Type}}) bool {
		return less(b, a)
	}
	return list.IsSortedWith(greater)
}
` + sortImplementation,
}
