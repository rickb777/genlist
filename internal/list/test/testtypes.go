package main

// +test List:"MapTo[Num1], MapTo[*Num2], With[Num1], With[Foo], SortWith"
type Thing struct {
	Name   string
	Number Num1
}

// methods where underlying type is ordered or numeric
// +test List:"MapTo[Foo]"
type Num1 Underlying

// +test * List
type Num2 Underlying

type Underlying int

// +test List:"MapTo[Num1]"
type Foo string

