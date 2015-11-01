package main

// methods where underlying type is ordered or numeric
// +test Option:"MapTo[Foo]"
type Num1 Underlying

type Underlying int

// +test Option
type Foo string

// +test * Option
type Bar int
