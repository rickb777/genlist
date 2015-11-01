package plumbing

const Plumbing = `
//-------------------------------------------------------------------------------------------------
// This plumbing suite provides standard functions for piping data between goroutines.
// All these functions run until the input channel is closed. They then close their output channel(s).
// None of these functions create a goroutine - this must be done at the call site.

// {{.TName}}Delta duplicates a stream of {{.PName}} to two output channels.
// When the sender closes the input channel, both output channels are closed then the function terminates.
func {{.TName}}Delta(in <-chan {{.PName}}, out1, out2 chan<- {{.PName}}) {
	for v := range in {
		select {
		case out1 <- v: out2 <- v
		case out2 <- v: out1 <- v
		}
	}
	close(out1)
	close(out2)
}

// {{.TName}}Zip interleaves two streams of {{.PName}}. Each input channel is used in turn, alternating between them.
// The function terminates when *both* input channels have been closed by their senders. The output channel is
// then closed also.
func {{.TName}}Zip(in1, in2 <-chan {{.PName}}, out chan<- {{.PName}}) {
	closed := false
	for v := range in1 {
		out <- v
		v, ok := <- in2
		if ok {
			out <- v
		} else {
			closed = true
		}
	}
	// need to drain in2 as well?
	if !closed {
		for v := range in1 {
		}
	}
	close(out)
}

// {{.TName}}BlackHole silently consumes a stream of {{.PName}}. It terminates when the sender closes the channel.
func {{.TName}}BlackHole(in <-chan {{.PName}}) {
	for v := range in {
		// om nom nom
	}
}

// {{.TName}}Filter filters a stream of {{.PName}}, silently dropping elements that do not match the predicate p.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}Filter(in <-chan {{.PName}}, out chan<- {{.PName}}, p func({{.PName}}) bool) {
	for v := range in {
		if p(v) {
			out <- v
		}
	}
	close(out)
}

// {{.TName}}Partition filters a stream of {{.PName}} into two output streams using a predicate p.
// When the sender closes the input channel, both output channels are closed then the function terminates.
func {{.TName}}Partition(in <-chan {{.PName}}, matching, others chan<- {{.PName}}, p func({{.PName}}) bool) {
	for v := range in {
		if p(v) {
			matching <- v
		} else {
			others <- v
		}
	}
	close(matching)
	close(others)
}

// {{.TName}}Map transforms a stream of {{.PName}} by applying a function to each item in the stream.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}Map(in <-chan {{.PName}}, out chan<- {{.PName}}, fn func({{.PName}}) {{.PName}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.TName}}FlatMap transforms a stream of {{.PName}} by applying a function to each item in the stream that
// gives zero or more results, all of which are sent out.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}FlatMap(in <-chan {{.PName}}, out chan<- {{.PName}}, fn func({{.PName}}) {{.TName}}Collection) {
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
