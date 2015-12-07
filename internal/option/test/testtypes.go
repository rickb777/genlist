package main

// methods where underlying type is ordered or numeric
// +test Option:"MapTo[Foo]"
type Num1 Underlying

type Underlying int

// +test Option
type Foo string

// +test * Option
type Bar int

// This contains a slice and should not be comparable
// +test Option
type WotHasSlice struct {
	Stuff []int
}

// This contains a map and should not be comparable
// +test Option
type WotHasMap struct {
	Stuff map[int]int
}

