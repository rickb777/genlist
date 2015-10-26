package main

// +test List:"MapTo[Num1], MapTo[*Num2], With[Num1], With[Colour], SortWith"
type Thing struct {
	Name   string
	Number Num1
}

type Colour string

// methods where underlying type is ordered or numeric
// +test List
type Num1 Underlying

// +test * List
type Num2 Underlying

type Underlying int
