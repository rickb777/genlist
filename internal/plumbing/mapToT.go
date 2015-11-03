package plumbing

const PlumbingMapToParamFunctions = `
// {{.TName}}MapTo{{.TypeParameter.Name}} transforms a stream of {{.PName}} to a stream of {{.TypeParameter}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}MapTo{{.TypeParameter.Name}}(in <-chan {{.PName}}, out chan<- {{.TypeParameter}},
		fn func({{.PName}}) {{.TypeParameter}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.TName}}FlatMapTo{{.TypeParameter.Name}} transforms a stream of {{.PName}} to a stream of {{.TypeParameter}}.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}FlatMapTo{{.TypeParameter.Name}}(in <-chan {{.PName}}, out chan<- {{.TypeParameter}},
		fn func({{.PName}}) {{.TypeParameter.Name}}Collection) {
	for vi := range in {
		c := fn(vi)
		if c.NonEmpty() {
			for vo := range c.Send() {
				out <- vo
			}
		}
	}
	close(out)
}

`
