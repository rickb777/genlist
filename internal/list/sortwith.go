package list

var SortWith = `
//-------------------------------------------------------------------------------------------------
// List:SortWith

//-----------------------------------------------------------------------------
// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.
//-----------------------------------------------------------------------------


// SortWith returns a new ordered {{.TName}}List, determined by a func defining ‘less’.
func (list {{.TName}}List) SortWith(less func({{.PName}}, {{.PName}}) bool) {{.TName}}List {
	result := make({{.TName}}List, len(list))
	copy(result, list)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSort{{.TName}}List(result, less, 0, n, maxDepth)
	return result
}

// IsSortedWith reports whether an instance of {{.TName}}List is sorted, using the pass func to define ‘less’.
func (list {{.TName}}List) IsSortedWith(less func({{.PName}}, {{.PName}}) bool) bool {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		if less(list[i], list[i-1]) {
			return false
		}
	}
	return true
}

// SortWithDesc returns a new, descending-ordered {{.TName}}List, determined by a func defining ‘less’.
func (list {{.TName}}List) SortWithDesc(less func({{.PName}}, {{.PName}}) bool) {{.TName}}List {
	greater := func(a, b {{.PName}}) bool {
		return less(b, a)
	}
	return list.SortWith(greater)
}

// IsSortedDesc reports whether an instance of {{.TName}}List is sorted in descending order, using the pass func to define ‘less’.
func (list {{.TName}}List) IsSortedWithDesc(less func({{.PName}}, {{.PName}}) bool) bool {
	greater := func(a, b {{.PName}}) bool {
		return less(b, a)
	}
	return list.IsSortedWith(greater)
}
` + sortImplementation
