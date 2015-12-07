package main

// +test List:"MapTo[Num1], MapTo[*Num2], MapTo[string], With[Num1], With[Foo], SortWith"
type Thing struct {
	Name   string
	Number Num1
}

// This contains a slice and should not be comparable
// +test List
type WotHasSlice struct {
	Stuff []int
}

// This contains a map and should not be comparable
// +test List
type WotHasMap struct {
	Stuff map[int]int
}

// methods where underlying type is ordered or numeric
// +test List:"MapTo[Foo]"
type Num1 Underlying

// +test * List
type Num2 Underlying

type Underlying int

// +test List:"MapTo[Num1]"
type Foo string

