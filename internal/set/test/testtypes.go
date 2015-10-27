package main

// +test Set:"MapTo[Num1]"
type Thing struct {
	Name   string
	Number Num1
}

// methods where underlying type is ordered or numeric
// +test Set:"MapTo[Foo]"
type Num1 Underlying

// +test * Set
//type Num2 Underlying

type Underlying int

// +test Set:"MapTo[Num1]"
type Foo string

