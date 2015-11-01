package list

const headTail = `

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// panics if list is empty
func (list {{.TName}}List) Head() {{.PName}} {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// panics if list is empty
func (list {{.TName}}List) Last() {{.PName}} {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// panics if list is empty
func (list {{.TName}}List) Tail() {{.TName}}Collection {
	return {{.TName}}List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// panics if list is empty
func (list {{.TName}}List) Init() {{.TName}}Collection {
	return {{.TName}}List(list[:len(list)-1])
}

// IsEmpty tests whether {{.TName}}List is empty.
func (list {{.TName}}List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether {{.TName}}List is empty.
func (list {{.TName}}List) NonEmpty() bool {
	return len(list) > 0
}

// IsSequence returns true for lists.
func (list {{.TName}}List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list {{.TName}}List) IsSet() bool {
	return false
}

// ToSlice gets all the list's elements in a plain slice. This is simply a type conversion.
func (list {{.TName}}List) ToSlice() []{{.PName}} {
	return []{{.PName}}(list)
}

{{if .Type.Underlying.IsBasic}}
// To{{.Type.Underlying.LongName}}s gets all the elements in a []{{.Type.Underlying}}.
func (list {{.TName}}List) To{{.Type.Underlying.LongName}}s() []{{.Type.Underlying}} {
	slice := make([]{{.Type.Underlying}}, len(list))
	for i, v := range list {
		slice[i] = {{.Type.Underlying}}(v)
	}
	return slice
}

{{end}}
// ToList simply returns the list in this case, but is part of the Collection interface.
func (list {{.TName}}List) ToList() {{.TName}}List {
	return list
}

{{if .Has.Set}}
// ToSet gets all the list's elements in a {{.TName}}Set.
func (list {{.TName}}List) ToSet() {{.TName}}Set {
	set := make(map[{{.TName}}]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	return {{.TName}}Set(set)
}

{{end}}
`
