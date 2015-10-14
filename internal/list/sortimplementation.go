package list

const sortImplementation = `

//-------------------------------------------------------------------------------------------------
// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swap{{.TName}}List(list {{.TName}}List, a, b int) {
	list[a], list[b] = list[b], list[a]
}

// Insertion sort
func insertionSort{{.TName}}List(list {{.TName}}List, less func({{.PName}}, {{.PName}}) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(list[j], list[j-1]); j-- {
			swap{{.TName}}List(list, j, j-1)
		}
	}
}

// siftDown implements the heap property on list[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown{{.TName}}List(list {{.TName}}List, less func({{.PName}}, {{.PName}}) bool, lo, hi, first int) {
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
		swap{{.TName}}List(list, first+root, first+child)
		root = child
	}
}

func heapSort{{.TName}}List(list {{.TName}}List, less func({{.PName}}, {{.PName}}) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown{{.TName}}List(list, less, i, hi, first)
	}

	// Pop elements, largest first, into end of list.
	for i := hi - 1; i >= 0; i-- {
		swap{{.TName}}List(list, first, first+i)
		siftDown{{.TName}}List(list, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values list[a], list[b], list[c] into list[a].
func medianOfThree{{.TName}}List(list {{.TName}}List, less func({{.PName}}, {{.PName}}) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(list[m1], list[m0]) {
		swap{{.TName}}List(list, m1, m0)
	}
	if less(list[m2], list[m1]) {
		swap{{.TName}}List(list, m2, m1)
	}
	if less(list[m1], list[m0]) {
		swap{{.TName}}List(list, m1, m0)
	}
	// now list[m0] <= list[m1] <= list[m2]
}

func swapRange{{.TName}}List(list {{.TName}}List, a, b, n int) {
	for i := 0; i < n; i++ {
		swap{{.TName}}List(list, a+i, b+i)
	}
}

func doPivot{{.TName}}List(list {{.TName}}List, less func({{.PName}}, {{.PName}}) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree{{.TName}}List(list, less, lo, lo+s, lo+2*s)
		medianOfThree{{.TName}}List(list, less, m, m-s, m+s)
		medianOfThree{{.TName}}List(list, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree{{.TName}}List(list, less, lo, m, hi-1)

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
				swap{{.TName}}List(list, a, b)
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
				swap{{.TName}}List(list, c-1, d-1)
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
		swap{{.TName}}List(list, b, c-1)
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
	swapRange{{.TName}}List(list, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRange{{.TName}}List(list, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSort{{.TName}}List(list {{.TName}}List, less func({{.PName}}, {{.PName}}) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSort{{.TName}}List(list, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot{{.TName}}List(list, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort{{.TName}}List(list, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSort{{.TName}}List(list, mhi, b)
		} else {
			quickSort{{.TName}}List(list, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSort{{.TName}}List(list, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSort{{.TName}}List(list, less, a, b)
	}
}
`
