package list

import "github.com/rickb777/typewriter"

var SortImplementation = &typewriter.Template{
	Name: "sortImplementation",
	Text: `

//-------------------------------------------------------------------------------------------------
// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swap{{.ListName}}(list {{.ListName}}, a, b int) {
	list[a], list[b] = list[b], list[a]
}

// Insertion sort
func insertionSort{{.ListName}}(list {{.ListName}}, less func({{.Type}}, {{.Type}}) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(list[j], list[j-1]); j-- {
			swap{{.ListName}}(list, j, j-1)
		}
	}
}

// siftDown implements the heap property on list[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown{{.ListName}}(list {{.ListName}}, less func({{.Type}}, {{.Type}}) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(list[first+child], list[first+child+1]) {
			child++
		}
		if !less(list[first+root], list[first+child]) {
			return
		}
		swap{{.ListName}}(list, first+root, first+child)
		root = child
	}
}

func heapSort{{.ListName}}(list {{.ListName}}, less func({{.Type}}, {{.Type}}) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown{{.ListName}}(list, less, i, hi, first)
	}

	// Pop elements, largest first, into end of list.
	for i := hi - 1; i >= 0; i-- {
		swap{{.ListName}}(list, first, first+i)
		siftDown{{.ListName}}(list, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values list[a], list[b], list[c] into list[a].
func medianOfThree{{.ListName}}(list {{.ListName}}, less func({{.Type}}, {{.Type}}) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(list[m1], list[m0]) {
		swap{{.ListName}}(list, m1, m0)
	}
	if less(list[m2], list[m1]) {
		swap{{.ListName}}(list, m2, m1)
	}
	if less(list[m1], list[m0]) {
		swap{{.ListName}}(list, m1, m0)
	}
	// now list[m0] <= list[m1] <= list[m2]
}

func swapRange{{.ListName}}(list {{.ListName}}, a, b, n int) {
	for i := 0; i < n; i++ {
		swap{{.ListName}}(list, a+i, b+i)
	}
}

func doPivot{{.ListName}}(list {{.ListName}}, less func({{.Type}}, {{.Type}}) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree{{.ListName}}(list, less, lo, lo+s, lo+2*s)
		medianOfThree{{.ListName}}(list, less, m, m-s, m+s)
		medianOfThree{{.ListName}}(list, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree{{.ListName}}(list, less, lo, m, hi-1)

	// Invariants are:
	//	list[lo] = pivot (set up by ChoosePivot)
	//	list[lo <= i < a] = pivot
	//	list[a <= i < b] < pivot
	//	list[b <= i < c] is unexamined
	//	list[c <= i < d] > pivot
	//	list[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(list[b], list[pivot]) { // list[b] < pivot
				b++
			} else if !less(list[pivot], list[b]) { // list[b] = pivot
				swap{{.ListName}}(list, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(list[pivot], list[c-1]) { // list[c-1] > pivot
				c--
			} else if !less(list[c-1], list[pivot]) { // list[c-1] = pivot
				swap{{.ListName}}(list, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// list[b] > pivot; list[c-1] < pivot
		swap{{.ListName}}(list, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRange{{.ListName}}(list, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRange{{.ListName}}(list, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSort{{.ListName}}(list {{.ListName}}, less func({{.Type}}, {{.Type}}) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSort{{.ListName}}(list, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot{{.ListName}}(list, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort{{.ListName}}(list, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSort{{.ListName}}(list, mhi, b)
		} else {
			quickSort{{.ListName}}(list, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSort{{.ListName}}(list, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSort{{.ListName}}(list, less, a, b)
	}
}
`}
