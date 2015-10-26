package main

// methods where underlying type is ordered or numeric
// +test Set
type Num1 Underlying

type Underlying int

// +test Set
type Foo string

// +test * Set
type Bar int

// +test Set
type Thing struct {
	Name   string
	Number Num1
}
