package main

// methods where underlying type is ordered or numeric
// +test Option:"With[Foo]"
type Other Underlying

type Underlying int

// +test Option
type Foo string

// +test * Option
type Bar int
