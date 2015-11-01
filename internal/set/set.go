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
// New{{.TName}}Set constructs a new set containing the supplied values, if any.
func New{{.TName}}Set(values ...{{.PName}}) {{.TName}}Set {
	set := make(map[{{.TName}}]struct{})
	for _, v := range values {
		set[v] = struct{}{}
	}
	return {{.TName}}Set(set)
}

{{if .Type.IsBasic}}
// New{{.TName}}SetFrom{{.Type.Underlying.LongName}}s constructs a new {{.TName}}Set from a []{{.Type.Underlying}}.
func New{{.TName}}SetFrom{{.Type.Underlying.LongName}}s(values []{{.Type.Underlying}}) {{.TName}}Set {
	set := make(map[{{.TName}}]struct{})
	for _, v := range values {
		set[{{.TName}}(v)] = struct{}{}
	}
	return {{.TName}}Set(set)
}

{{end}}
// Build{{.TName}}SetFrom constructs a new {{.TName}}Set from a channel that supplies values
// until it is closed.
func Build{{.TName}}SetFrom(source <-chan {{.PName}}) {{.TName}}Set {
	set := make(map[{{.TName}}]struct{})
	for v := range source {
		set[v] = struct{}{}
	}
	return {{.TName}}Set(set)
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns false for sets.
func (set {{.TName}}Set) IsSequence() bool {
	return false
}

// IsSet returns true for sets.
func (set {{.TName}}Set) IsSet() bool {
	return true
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

// Head gets an arbitrary element.
func (set {{.TName}}Set) Head() {{.PName}} {
	for v := range set {
		return v
	}
	panic("Set is empty")
}

// ToSlice gets all the set's elements in a plain slice.
func (set {{.TName}}Set) ToSlice() []{{.PName}} {
	slice := make([]{{.PName}}, set.Size())
	i := 0
	for v := range set {
		slice[i] = v
		i++
	}
	return slice
}

{{if .Type.Underlying.IsBasic}}
// To{{.Type.Underlying.LongName}}s gets all the elements in a []{{.Type.Underlying}}.
func (set {{.TName}}Set) To{{.Type.Underlying.LongName}}s() []{{.Type.Underlying}} {
	slice := make([]{{.Type.Underlying}}, len(set))
	i := 0
	for v := range set {
		slice[i] = {{.Type.Underlying}}(v)
		i++
	}
	return slice
}

{{end}}
{{if .Has.List}}
// ToList gets all the set's elements in a in {{.Name}}List.
func (set {{.TName}}Set) ToList() {{.TName}}List {
	return {{.TName}}List(set.ToSlice())
}

{{end}}
// ToSet gets the current set, which requires no further conversion.
func (set {{.TName}}Set) ToSet() {{.TName}}Set {
	return set
}


` + setAlgebra + addRemoveFunctions + iterationFunctions + predicatedFunctions +
numericFunctions + orderedFunctions + mkstringFunctions
