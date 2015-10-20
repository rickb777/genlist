package set

import "github.com/rickb777/golist/internal/sequence"

const Set = sequence.Collection + `
//-------------------------------------------------------------------------------------------------
// {{.TName}}Set is a typesafe set of {{.TName}} items. The implementation is based on Go maps.

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
`
