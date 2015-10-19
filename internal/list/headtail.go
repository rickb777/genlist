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
func (list {{.TName}}List) Tail() {{.TName}}Seq {
	return {{.TName}}List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// panics if list is empty
func (list {{.TName}}List) Init() {{.TName}}Seq {
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

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list {{.TName}}List) ToList() {{.TName}}List {
	return list
}
`
