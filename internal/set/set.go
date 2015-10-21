package set

import "github.com/rickb777/golist/internal/sequence"

const Set = sequence.Collection + `
//-------------------------------------------------------------------------------------------------
// {{.TName}}Set is a typesafe set of {{.TName}} items. Instances are essentially immutable.
// The set-agebra functions Union, Intersection and Difference allow new variants to be constructed
// easily.
//
// The implementation is based on Go maps.

type {{.TName}}Set map[{{.TName}}]struct{}

//-------------------------------------------------------------------------------------------------

func New{{.TName}}Set(e ...{{.TName}}) {{.TName}}Set {
	set := make(map[{{.TName}}]struct{})
	for _, v := range e {
		set[v] = struct{}{}
	}
	return {{.TName}}Set(set)
}

func (set {{.TName}}Set) Size() int {
	return len(set)
}

func (set {{.TName}}Set) IsEmpty() bool {
	return len(set) == 0
}

func (set {{.TName}}Set) NonEmpty() bool {
	return len(set) > 0
}

// ToSlice gets all the elements in {{.Name}}Set in a slice.
func (set {{.TName}}Set) ToSlice() []{{.TName}} {
	slice := make([]{{.TName}}, 0, len(set))
	for v := range set {
		slice = append(slice, v)
	}
	return slice
}

` + setAlgebra + addRemoveFunctions + iterationFunctions + predicatedFunctions + mkstringFunctions
