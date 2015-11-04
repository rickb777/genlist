package golist

type flags struct {
	Collection bool
	List       bool
	Option     bool
	Set        bool
	Plumbing   bool
	Tag        map[string]bool
}

func newFlags() flags {
	flags := flags{}
	flags.Tag = make(map[string]bool)
	return flags
}

func (f flags) setFlag(name string) flags {
	switch name {
	case listName:   f.List = true
	case optionName: f.Option = true
	case setName:    f.Set = true
	default:         f.Tag[name] = true
	}
	return f
}

