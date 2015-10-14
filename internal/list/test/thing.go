package main

// +test List:"MapTo[Num1], MapTo[*Num2], With[Num1], With[Colour], SortWith"
type Thing struct {
	Name   string
	Number Num1
}

type Colour string
