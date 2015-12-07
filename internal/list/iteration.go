package list

const iterationFunctions = `
//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of {{.TName}}List return true for the passed func.
func (list {{.TName}}List) Exists(fn func({{.PName}}) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of {{.TName}}List return true for the passed func.
func (list {{.TName}}List) Forall(fn func({{.PName}}) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over {{.TName}}List and executes the passed func against each element.
func (list {{.TName}}List) Foreach(fn func({{.PName}})) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
func (list {{.TName}}List) Send() <-chan {{.PName}} {
	ch := make(chan {{.PName}})
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of {{.TName}}List with all elements in the reverse order.
func (list {{.TName}}List) Reverse() {{.TName}}List {
    numItems := len(list)
    result := make({{.TName}}List, numItems)
    last := numItems - 1
    for i, v := range list {
		result[last - i] = v
    }
    return result
}

// Shuffle returns a shuffled copy of {{.TName}}List, using a version of the Fisher-Yates shuffle.
func (list {{.TName}}List) Shuffle() {{.TName}}List {
    numItems := len(list)
    result := make({{.TName}}List, numItems)
    copy(result, list)
    for i := 0; i < numItems; i++ {
        r := i + rand.Intn(numItems-i)
        result.Swap(i, r)
    }
    return result
}

`
