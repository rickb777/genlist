package slice

import "github.com/clipperhouse/typewriter"

var size = &typewriter.Template{
	Name: "Size",
	Text: `
// Size gets the length of {{.SliceName}}.
func (rcv {{.SliceName}}) Size() int {
	return len(rcv)
}

// IsEmpty tests whether {{.SliceName}} is empty.
func (rcv {{.SliceName}}) IsEmpty() bool {
	return len(rcv) == 0
}

// NonEmpty tests whether {{.SliceName}} is empty.
func (rcv {{.SliceName}}) NonEmpty() bool {
	return len(rcv) > 0
}
`}
