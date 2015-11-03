package main

// +test Set:"MapTo[Num1], MapTo[string], With[Num1], With[Foo]"
//type Thing struct {
//	Name   string
//	Number Num1
//}

// methods where underlying type is ordered or numeric
// +test List:"MapTo[Foo] Option Set Plumbing"
type Num1 Underlying

// +test * Set
//type Num2 Underlying

type Underlying int

// +test List:"MapTo[Foo] Option Set"
type Foo string

