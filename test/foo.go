package main

// +test List:" Option Set MapTo[Foo2]"
type Foo1 Underlying

// +test Option:" List Set MapTo[Foo1]"
type Foo2 Underlying

// +test * List:" Option MapTo[Foo4]"
type Foo3 Underlying

// +test * Option:" List MapTo[Foo3]"
type Foo4 Underlying

type Underlying int
