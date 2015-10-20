package set

import "github.com/rickb777/golist/internal/sequence"

const Set = sequence.Sequence + `
//-------------------------------------------------------------------------------------------------
// {{.TName}}Set is a typesafe set of {{.TName}} items. The implementation is based on Go maps.

type {{.TName}}Set map[{{.TName}}]struct{}

//-------------------------------------------------------------------------------------------------

func New{{.TName}}Set(e {{.TName}}...) {
	set := make(map[{{.TName}}]s
}
`
