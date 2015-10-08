package main

// +test slice:"Comparable, Comparable[Other], Ordered[Other], Numeric[Other], Aggregate[Other], MapTo[Other], First, SortWith"
type Thing struct {
	Name   string
	Number Other
}
